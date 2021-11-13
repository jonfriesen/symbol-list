package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jonfriesen/symbol-list/internal/export"
	"github.com/jonfriesen/symbol-list/internal/model"
	"github.com/jonfriesen/symbol-list/internal/sources/crypto"
	"github.com/jonfriesen/symbol-list/internal/sources/nasdaq"
	"github.com/jonfriesen/symbol-list/internal/sources/tsx"
	"golang.org/x/sync/errgroup"
)

func main() {

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

	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		log.Fatalln("failed to create directory", err)
	}

	fName := time.Now().Format("2006-01-02")

	col := &model.SecurityExport{
		Date:       fName,
		Securities: securities,
	}

	err = export.JSON("data/"+fName, col)
	if err != nil {
		log.Fatalln(err)
	}

	err = export.CSV("data/"+fName, col)
	if err != nil {
		log.Fatalln(err)
	}

	cryptoCol := &model.CryptoExport{
		Date:       fName,
		Currencies: cryptoCurrencies,
	}

	err = export.JSON("data/"+fName+"-crypto", cryptoCol)
	if err != nil {
		log.Fatalln(err)
	}

	err = export.CSV("data/"+fName+"-crypto", cryptoCol)
	if err != nil {
		log.Fatalln(err)
	}
}
