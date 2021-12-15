package models

type Config struct {
	CurrencyURL   string
	CurrencyToken string
}

func NewConfig() Config {
	return Config{
		CurrencyURL:   "https://api.currencylayer.com/live?access_key=%s",
		CurrencyToken: "4541fb80fdd854f1c12975a297747ab5",
	}
}
