package application

import (
	"context"

	"msbeer.com/src/adapter"
	"msbeer.com/src/infra"
	"msbeer.com/src/models"
)

//BeerApplication Interface
type BeerApplication interface {
	SearchBeers(ctx context.Context) ([]models.BeerItem, error)
	AddBeers(ctx context.Context, beerItem models.BeerItem) error
	SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error)
	BoxBeerPriceById(ctx context.Context, ID int, quantity int, currency string) (*models.BeerBox, error)
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
func (a BeerApplicationImpl) SearchBeers(ctx context.Context) ([]models.BeerItem, error) {
	return a.Infrastructure.SearchBeers(ctx)
}

//AddBeers adds a brand new beer into db
func (a BeerApplicationImpl) AddBeers(ctx context.Context, beerItem models.BeerItem) error {
	return a.Infrastructure.AddBeers(ctx, beerItem)
}

//SearchBeerById searches a beer by ID
func (a BeerApplicationImpl) SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error) {
	return a.Infrastructure.SearchBeerById(ctx, ID)
}

//BoxBeerPriceById searches for box total value by beer ID
func (a BeerApplicationImpl) BoxBeerPriceById(ctx context.Context, ID int, quantity int, currency string) (*models.BeerBox, error) {
	currencyInfo, err := a.Adapter.GetCurrencyInfo()
	if err != nil {
		return nil, err
	}

	currencyValue, err := currencyInfo.GetCurrentValue(currency)
	if err != nil {
		return nil, err
	}

	beerItem, err := a.Infrastructure.SearchBeerById(ctx, ID)
	if err != nil {
		return nil, err
	}

	totalAmount := (beerItem.Price * float64(quantity)) * currencyValue

	return &models.BeerBox{
		PriceTotal: totalAmount,
	}, nil
}
