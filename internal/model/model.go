package model

type Security struct {
	Symbol      string       `json:"symbol,omitempty"`
	Name        string       `json:"name,omitempty"`
	Exchange    string       `json:"exchange,omitempty"`
	YahooSymbol string       `json:"yahoo_symbol,omitempty"`
	Instruments []Instrument `json:"instruments,omitempty"`
}

type Instrument struct {
	Name   string `json:"name,omitempty"`
	Symbol string `json:"symbol,omitempty"`
}
