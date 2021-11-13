package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/jonfriesen/symbol-list/internal/export"
	"github.com/jonfriesen/symbol-list/internal/model"
	"github.com/jonfriesen/symbol-list/internal/sources/nasdaq"
	"github.com/jonfriesen/symbol-list/internal/sources/tsx"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	eg := errgroup.Group{}

	var securitiesMutext sync.Mutex
	var securities []*model.Security

	eg.Go(func() error {
		fmt.Println("Retrieving Nasdaq listed securities")
		nasdaqClient := nasdaq.New()
		ls, err := nasdaqClient.GetListedSymbols()
		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
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
			fmt.Println(err)
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
			fmt.Println(err)
			return err
		}

		securitiesMutext.Lock()
		securities = append(securities, tsxv...)
		securitiesMutext.Unlock()

		return nil
	})

	err := eg.Wait()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Found %d securities\n", len(securities))

	err = os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Println(errors.Wrap(err, "failed to create directory"))
		return
	}

	fName := time.Now().Format("2006-01-02")

	col := &model.Export{
		Date: fName,
		Data: securities,
	}

	err = export.JSON("data/"+fName, col)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = export.CSV("data/"+fName, col)
	if err != nil {
		fmt.Println(err)
		return
	}
}
