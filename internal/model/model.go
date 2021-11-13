package model

type SecurityExport struct {
	Date       string      `json:"date,omitempty"`
	Securities []*Security `json:"data,omitempty"`
}

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

func (s *SecurityExport) Data() [][]string {
	data := make([][]string, len(s.Securities))

	for i, s := range s.Securities {
		data[i] = []string{s.Symbol, s.Name, s.Exchange, s.YahooSymbol}
		for _, y := range s.Instruments {
			data[i] = append(data[i], y.Name)
			data[i] = append(data[i], y.Symbol)
		}
	}

	return data
}

func (e SecurityExport) CSVHeader() []string {
	h := []string{"Symbol", "Name", "Exchange", "Yahoo Symbol"}

	instruCount := 0
	for _, s := range e.Securities {
		if len(s.Instruments) > instruCount {
			instruCount = len(s.Instruments)
		}
	}

	for i := 0; i < instruCount; i++ {
		h = append(h, "Instrument Name")
		h = append(h, "Instrument Symbol")
	}

	return h
}

type CryptoExport struct {
	Date       string    `json:"date,omitempty"`
	Currencies []*Crypto `json:"data,omitempty"`
}

type Crypto struct {
	Symbol  string `json:"symbol,omitempty"`
	Name    string `json:"name,omitempty"`
	BuiltOn string `json:"built_on,omitempty"`
}

func (s *CryptoExport) Data() [][]string {
	data := make([][]string, len(s.Currencies))

	for i, s := range s.Currencies {
		data[i] = []string{s.Symbol, s.Name, s.BuiltOn}
	}

	return data
}

func (e CryptoExport) CSVHeader() []string {
	return []string{"Symbol", "Name", "BuiltOn"}
}
