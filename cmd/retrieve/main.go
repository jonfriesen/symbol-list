package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"

	"github.com/jonfriesen/symbol-list/internal/export"
	"github.com/jonfriesen/symbol-list/internal/importer"
	"github.com/jonfriesen/symbol-list/internal/model"
	"github.com/jonfriesen/symbol-list/internal/sources/crypto"
	"github.com/jonfriesen/symbol-list/internal/sources/nasdaq"
	"github.com/jonfriesen/symbol-list/internal/sources/tsx"
	"golang.org/x/sync/errgroup"
)

const securityPattern = `^\d{4}-\d{2}-\d{2}.json$`
const cryptoPattern = `^\d{4}-\d{2}-\d{2}-crypto.json$`

func main() {

	saveDir := flag.String("dir", "data/", "Directory to save data files.")
	flag.Parse()

	eg := errgroup.Group{}

	var securitiesMutext sync.Mutex
	var securities []*model.Security

	var cryptoMutext sync.Mutex
	var cryptoCurrencies []*model.Crypto

	eg.Go(func() error {
		fmt.Println("Retrieving Nasdaq listed securities")
		nasdaqClient := nasdaq.New()
		ls, err := nasdaqClient.GetListedSymbols()
		if err != nil {
			return err
		}

		securitiesMutext.Lock()
		securities = append(securities, ls...)
		securitiesMutext.Unlock()

		return nil
	})

	eg.Go(func() error {
		fmt.Println("Retrieving Nasdaq other securities")
		nasdaqClient := nasdaq.New()
		os, err := nasdaqClient.GetOtherSymbols()
		if err != nil {
			return err
		}

		securitiesMutext.Lock()
		securities = append(securities, os...)
		securitiesMutext.Unlock()

		return nil
	})

	eg.Go(func() error {
		fmt.Println("Retrieving TSX securities")
		tsxClient := tsx.New()
		tsx, err := tsxClient.GetSymbols()
		if err != nil {
			return err
		}

		securitiesMutext.Lock()
		securities = append(securities, tsx...)
		securitiesMutext.Unlock()

		return nil
	})

	eg.Go(func() error {
		fmt.Println("Retrieving TSXV securities")
		tsxClient := tsx.New()
		tsxv, err := tsxClient.GetVentureSymbols()
		if err != nil {
			return err
		}

		securitiesMutext.Lock()
		securities = append(securities, tsxv...)
		securitiesMutext.Unlock()

		return nil
	})

	eg.Go(func() error {
		fmt.Println("Retrieving Cryptocurrencies")
		cryptoClient := crypto.New()
		c, err := cryptoClient.GetSymbols()
		if err != nil {
			return err
		}

		cryptoMutext.Lock()
		cryptoCurrencies = append(cryptoCurrencies, c...)
		cryptoMutext.Unlock()

		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Found %d securities\n", len(securities))
	fmt.Printf("Found %d cryptocurrencies\n", len(cryptoCurrencies))

	// get current latest diff files
	oldSecurity, err := importer.OrderedList(securityPattern, 1, *saveDir)
	if err != nil {
		log.Fatalln(err)
	}
	oldCrypto, err := importer.OrderedList(cryptoPattern, 1, *saveDir)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.MkdirAll(*saveDir, os.ModePerm)
	if err != nil {
		log.Fatalln("failed to create directory", err)
	}

	fName := time.Now().Format("2006-01-02")

	col := &model.SecurityExport{
		Date:       fName,
		Securities: securities,
	}

	err = export.JSON(path.Join(*saveDir, fName), col)
	if err != nil {
		log.Fatalln(err)
	}

	err = export.CSV(path.Join(*saveDir, fName), col)
	if err != nil {
		log.Fatalln(err)
	}

	cryptoCol := &model.CryptoExport{
		Date:       fName,
		Currencies: cryptoCurrencies,
	}

	err = export.JSON(path.Join(*saveDir, fName)+"-crypto", cryptoCol)
	if err != nil {
		log.Fatalln(err)
	}

	err = export.CSV(path.Join(*saveDir, fName)+"-crypto", cryptoCol)
	if err != nil {
		log.Fatalln(err)
	}

	// start diff creation
	fmt.Println("oldSecuirty", len(oldSecurity))
	if len(oldSecurity) > 0 {
		oldSec := &model.SecurityExport{}
		err := importer.JSON(oldSecurity[0], oldSec)
		if err != nil {
			log.Fatalln(err)
		}

		securitiesDiff := &model.SecurityExportDiff{}
		err = securitiesDiff.Diff(col.Securities, oldSec.Securities)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Securities Added %d\n", len(securitiesDiff.Added))
		fmt.Printf("Securities Removed %d\n", len(securitiesDiff.Removed))

		err = export.JSON(path.Join(*saveDir, fName)+"-diff", securitiesDiff)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("oldCrypto", len(oldCrypto))
	if len(oldCrypto) > 0 {
		oldCry := &model.CryptoExport{}
		err := importer.JSON(oldCrypto[0], oldCry)
		if err != nil {
			log.Fatalln(err)
		}

		cryptoDiff := &model.CryptoExportDiff{}
		err = cryptoDiff.Diff(cryptoCol.Currencies, oldCry.Currencies)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Crypto Added %d\n", len(cryptoDiff.Added))
		fmt.Printf("Crypto Removed %d\n", len(cryptoDiff.Removed))

		err = export.JSON(path.Join(*saveDir, fName)+"-diff-crypto", cryptoDiff)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
