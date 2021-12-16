package models

type Config struct {
	CurrencyURL        string
	CurrencyToken      string
	DbConnectionString string
	DbProvider         string
}

func NewConfig() Config {
	return Config{
		CurrencyURL:        "https://api.currencylayer.com/live?access_key=%s",
		CurrencyToken:      "4541fb80fdd854f1c12975a297747ab5",
		DbConnectionString: "server=PAVILLALOBOS;user id=;trusted_connection=true;database=msbeer;app name=msbeer",
		DbProvider:         "mssql",
	}
}
