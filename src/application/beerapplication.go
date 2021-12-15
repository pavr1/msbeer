package application

import (
	"msbeer.com/src/adapter"
	"msbeer.com/src/infra"
	"msbeer.com/src/models"
)

//BeerApplication Interface
type BeerApplication interface {
	SearchBeers() ([]models.BeerItem, error)
	AddBeers(id int, name, brewery, country string, price float64, currency string) error
	SearchBeerById(ID int) (models.BeerItem, error)
	BoxBeerPriceById(ID int, quantity int, currency string) (*models.BeerBox, error)
}

//BeerApplicationImpl structure
type BeerApplicationImpl struct {
	Adapter        adapter.BeerAdapter
	Infrastructure infra.BeerInfra
}

//NewApplication creates an instance of BeerApplicationImpl
func NewApplication(adapter adapter.BeerAdapter, infrastructure infra.BeerInfra) BeerApplication {
	return BeerApplicationImpl{
		Adapter:        adapter,
		Infrastructure: infrastructure,
	}
}

//SearchBeers searches all list of beers existent in db
func (a BeerApplicationImpl) SearchBeers() ([]models.BeerItem, error) {
	return a.Infrastructure.SearchBeers()
}

//AddBeers adds a brand new beer into db
func (a BeerApplicationImpl) AddBeers(id int, name, brewery, country string, price float64, currency string) error {
	beerItem := models.BeerItem{
		ID:       id,
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}

	return a.Infrastructure.AddBeers(beerItem)
}

//SearchBeerById searches a beer by ID
func (a BeerApplicationImpl) SearchBeerById(ID int) (models.BeerItem, error) {
	return a.Infrastructure.SearchBeerById(ID)
}

//BoxBeerPriceById searches for box total value by beer ID
func (a BeerApplicationImpl) BoxBeerPriceById(ID int, quantity int, currency string) (*models.BeerBox, error) {
	currencyInfo, err := a.Adapter.GetCurrencyInfo()
	if err != nil {
		return nil, err
	}

	currencyValue, err := currencyInfo.GetCurrentValue(currency)
	if err != nil {
		return nil, err
	}

	beerItem, err := a.Infrastructure.SearchBeerById(ID)
	if err != nil {
		return nil, err
	}

	totalAmount := (beerItem.Price * float64(quantity)) * currencyValue

	return &models.BeerBox{
		PriceTotal: totalAmount,
	}, nil
}
