package model

type Export struct {
	Date string      `json:"date,omitempty"`
	Data []*Security `json:"data,omitempty"`
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

func (s *Security) Row() []string {
	r := []string{s.Symbol, s.Name, s.Exchange, s.YahooSymbol}

	for _, i := range s.Instruments {
		r = append(r, i.Name)
		r = append(r, i.Symbol)
	}

	return r
}

func (e Export) CSVHeader() []string {
	h := []string{"Symbol", "Name", "Exchange", "Yahoo Symbol"}

	instruCount := 0
	for _, s := range e.Data {
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
