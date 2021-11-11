package tsx

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/jonfriesen/symbol-list/internal/model"
)

const tsxURL = "https://www.tsx.com/json/company-directory/search/tsx/^*"
const tsxvURL = "https://www.tsx.com/json/company-directory/search/tsxv/^*"

type client struct {
	httpClient *http.Client
}

type tsxResponse struct {
	Results []*model.Security `json:"results"`
}

func New() *client {
	timeout := time.Duration(30) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	httpClient := &http.Client{
		Transport: transport,
	}

	c := &client{
		httpClient: httpClient,
	}

	return c
}

func (c *client) GetSymbols() ([]*model.Security, error) {
	resp, err := c.httpClient.Get(tsxURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	symbols, err := parseTSXSymbols(resp.Body)
	if err != nil {
		return nil, err
	}

	for _, s := range symbols {
		s := s
		s.Exchange = "TSX"
		s.YahooSymbol = s.Symbol + ".TO"
	}

	return symbols, nil
}

func (c *client) GetVentureSymbols() ([]*model.Security, error) {
	resp, err := c.httpClient.Get(tsxvURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	symbols, err := parseTSXSymbols(resp.Body)
	if err != nil {
		return nil, err
	}

	for _, s := range symbols {
		s := s
		s.Exchange = "TSXV"
		s.YahooSymbol = s.Symbol + ".V"
	}

	return symbols, nil
}

func parseTSXSymbols(body io.ReadCloser) ([]*model.Security, error) {
	var tsxResp tsxResponse
	err := json.NewDecoder(body).Decode(&tsxResp)
	if err != nil {
		return nil, err
	}

	return tsxResp.Results, nil
}
