package adapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"msbeer.com/src/models"
)

type BeerAdapter interface {
	GetCurrencyInfo() (*models.Currency, error)
}

type BeerAdapterImpl struct {
	CurrencyApiUrl string
	HttpClient     *http.Client
}

func NewBeerAdapter(currencyApiUrl string, httpClient *http.Client) BeerAdapter {
	return BeerAdapterImpl{
		CurrencyApiUrl: currencyApiUrl,
		HttpClient:     httpClient,
	}
}

func (a BeerAdapterImpl) GetCurrencyInfo() (*models.Currency, error) {
	resp, err := a.HttpClient.Get(a.CurrencyApiUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	currency := models.Currency{}
	err = json.Unmarshal(body, &currency)
	if err != nil {
		return nil, err
	}

	return &currency, nil
}
