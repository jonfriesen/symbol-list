package nasdaq

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/jonfriesen/symbol-list/internal/model"
)

const nasdaqFTP = "ftp.nasdaqtrader.com:21"
const nasdaqListed = "SymbolDirectory/nasdaqlisted.txt"
const otherListed = "SymbolDirectory/otherlisted.txt"
const timeout = time.Duration(30) * time.Second

type client struct {
	conn *ftp.ServerConn
}

// New creates a new Nasdaq client.
func New() *client {
	f, err := ftp.Dial(nasdaqFTP, ftp.DialWithTimeout(timeout))
	if err != nil {
		log.Fatal(err)
	}
	err = f.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	return &client{
		conn: f,
	}
}

func (c *client) GetListedSymbols() ([]*model.Security, error) {
	var securities []*model.Security

	r, err := c.conn.Retr(nasdaqListed)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	securities, err = parseListedSymbols(r)
	if err != nil {
		return nil, err
	}

	return securities, nil
}

func parseListedSymbols(r io.Reader) ([]*model.Security, error) {
	var securities []*model.Security
	csvr := csv.NewReader(r)
	csvr.Comma = '|'

	// skip header
	if _, err := csvr.Read(); err != nil {
		return nil, err
	}

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return nil, err
		}

		if strings.HasPrefix(row[0], "File Creation Time") || row[3] == "Y" {
			continue
		}

		securities = append(securities, &model.Security{
			Symbol:      row[0],
			Name:        row[1],
			Exchange:    getListedExchange(row[2]),
			YahooSymbol: getYahooSymbol(row[0]),
		})
	}

	return securities, nil
}

func getListedExchange(indicator string) string {
	switch indicator {
	case "Q":
		return "NASDAQ Global Select MarketSM"
	case "G":
		return "NASDAQ Global MarketSM"
	case "S":
		return "NASDAQ Capital Market"
	}
	return "NASDAQ"
}

func (c *client) GetOtherSymbols() ([]*model.Security, error) {
	var securities []*model.Security

	r, err := c.conn.Retr(otherListed)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	securities, err = parseOtherSymbols(r)
	if err != nil {
		return nil, err
	}

	return securities, nil
}

func parseOtherSymbols(r io.Reader) ([]*model.Security, error) {
	var securities []*model.Security
	csvr := csv.NewReader(r)
	csvr.Comma = '|'

	// skip header
	if _, err := csvr.Read(); err != nil {
		return nil, err
	}

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF || strings.HasPrefix(row[0], "File Creation Time") {
				err = nil
				break
			}
			fmt.Printf("%v\n", row)
			return nil, err
		}

		// skip test symbols
		if row[4] == "Y" {
			continue
		}

		securities = append(securities, &model.Security{
			Symbol:      row[0],
			Name:        row[1],
			Exchange:    getOtherExchange(row[2]),
			YahooSymbol: getYahooSymbol(row[0]),
		})
	}

	return securities, nil
}

func getOtherExchange(indicator string) string {
	switch indicator {
	case "A":
		return "NYSE MKT"
	case "N":
		return "NYSE"
	case "P":
		return "NYSE ARCA"
	case "Z":
		return "BATS"
	case "V":
		return "IEXG"
	}
	return "UNKNOWN"
}

func getYahooSymbol(t string) string {
	if n := strings.Count(t, "."); n > 1 {
		// if the ticker has multiple dots we create a yahoo formated version
		// and append it to the map so we can detect both versions and
		// translate them to the same ticker later
		yfinT := strings.Replace(t, ".", "-", n-1)
		return yfinT
	}

	return t
}
