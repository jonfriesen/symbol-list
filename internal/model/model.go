package model

type SecurityExport struct {
	Date       string      `json:"date,omitempty"`
	Securities []*Security `json:"data,omitempty"`
}
type SecurityExportDiff struct {
	Date    string      `json:"date,omitempty"`
	Added   []*Security `json:"added,omitempty"`
	Removed []*Security `json:"removed,omitempty"`
}

type Security struct {
	Symbol      string       `json:"symbol,omitempty"`
	Name        string       `json:"name,omitempty"`
	Exchange    string       `json:"exchange,omitempty"`
	YahooSymbol string       `json:"yahoo_symbol,omitempty"`
	Instruments []Instrument `json:"instruments,omitempty"`
}

func (s Security) Signature() string {
	return s.Symbol + s.Exchange
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

func (s *SecurityExportDiff) Diff(newSecurities []*Security, oldSecurities []*Security) error {
	s.Added = []*Security{}
	s.Removed = []*Security{}

	// create a map of old securities for efficient searching
	oldSecuritiesMap := make(map[string]*Security)
	for _, oldSecurity := range oldSecurities {
		oldSecuritiesMap[oldSecurity.Signature()] = oldSecurity
	}

	// iterate through the new securities and compare with the old securities
	for _, newSecurity := range newSecurities {
		if _, ok := oldSecuritiesMap[newSecurity.Signature()]; ok {
			// the security already exists, remove it from the map so that
			// at the end, the remaining securities in the map are the removed ones
			delete(oldSecuritiesMap, newSecurity.Signature())
		} else {
			// the security is new, add it to the added list
			s.Added = append(s.Added, newSecurity)
		}
	}

	// the remaining securities in the map are the removed ones
	for _, oldSecurity := range oldSecuritiesMap {
		s.Removed = append(s.Removed, oldSecurity)
	}

	return nil
}

type CryptoExport struct {
	Date       string    `json:"date,omitempty"`
	Currencies []*Crypto `json:"data,omitempty"`
}

type CryptoExportDiff struct {
	Date    string    `json:"date,omitempty"`
	Added   []*Crypto `json:"added,omitempty"`
	Removed []*Crypto `json:"removed,omitempty"`
}

type Crypto struct {
	Symbol  string `json:"symbol,omitempty"`
	Name    string `json:"name,omitempty"`
	BuiltOn string `json:"built_on,omitempty"`
}

func (c *Crypto) Signature() string {
	return c.Symbol + c.BuiltOn
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

func (s *CryptoExportDiff) Diff(new []*Crypto, old []*Crypto) error {
	s.Added = []*Crypto{}
	s.Removed = []*Crypto{}

	// create a map of old crypto for efficient searching
	oldMap := make(map[string]*Crypto)
	for _, oldCrypto := range old {
		oldMap[oldCrypto.Signature()] = oldCrypto
	}

	// iterate through the new crypto and compare with the old crypto
	for _, newCrypto := range new {
		if _, ok := oldMap[newCrypto.Signature()]; ok {
			// the crypto already exists, remove it from the map so that
			// at the end, the remaining crypto in the map are the removed ones
			delete(oldMap, newCrypto.Signature())
		} else {
			// the crypto is new, add it to the added list
			s.Added = append(s.Added, newCrypto)
		}
	}

	// the remaining crypto in the map are the removed ones
	for _, oldCrypto := range oldMap {
		s.Removed = append(s.Removed, oldCrypto)
	}

	return nil
}
