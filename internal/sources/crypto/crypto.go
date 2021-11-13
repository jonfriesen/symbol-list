package crypto

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/jonfriesen/symbol-list/internal/model"
	"github.com/pkg/errors"
)

const cryptoCompareAPIURL = "https://min-api.cryptocompare.com/data/all/coinlist"

type client struct {
	httpClient *http.Client
}

type coinCompareResp struct {
	Data map[string]*struct {
		Name    string `json:"FullName"`
		BuiltOn string `json:"BuiltOn"`
	} `json:"Data"`
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

func (c *client) GetSymbols() ([]*model.Crypto, error) {
	resp, err := c.httpClient.Get(cryptoCompareAPIURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ccResp coinCompareResp
	err = json.NewDecoder(resp.Body).Decode(&ccResp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode coin compare response")
	}

	symbols := make([]*model.Crypto, 0, len(ccResp.Data))
	for k, v := range ccResp.Data {
		symbols = append(symbols, &model.Crypto{
			Symbol:  k,
			Name:    v.Name,
			BuiltOn: v.BuiltOn,
		})
	}

	return symbols, nil
}
