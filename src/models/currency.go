package models

import "errors"

type Currency struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int64  `json:"timestamp"`
	Source    string `json:"souce"`
	Quote     Quote  `json:"quote"`
}

type Quote struct {
	USDAED float64 `json:"USDAED"`
	USDAFN float64 `json:"USDAFN"`
	USDALL float64 `json:"USDALL"`
	USDAMD float64 `json:"USDAMD"`
	USDANG float64 `json:"USDANG"`
	USDAOA float64 `json:"USDAOA"`
	USDARS float64 `json:"USDARS"`
	USDAUD float64 `json:"USDAUD"`
	USDAWG float64 `json:"USDAWG"`
	USDAZN float64 `json:"USDAZN"`
	USDBAM float64 `json:"USDBAM"`
}

func (c Currency) GetCurrentValue(currency string) (float64, error) {
	switch currency {
	case "USDAED":
		return c.Quote.USDAED, nil
	case "USDAFN":
		return c.Quote.USDAFN, nil
	case "USDALL":
		return c.Quote.USDALL, nil
	case "USDAMD":
		return c.Quote.USDAMD, nil
	case "USDANG":
		return c.Quote.USDANG, nil
	case "USDAOA":
		return c.Quote.USDAOA, nil
	case "USDARS":
		return c.Quote.USDARS, nil
	case "USDAUD":
		return c.Quote.USDAUD, nil
	case "USDAWG":
		return c.Quote.USDAWG, nil
	case "USDAZN":
		return c.Quote.USDAZN, nil
	case "USDBAM":
		return c.Quote.USDBAM, nil
	default:
		return -1, errors.New("unsupported currency")
	}
}
