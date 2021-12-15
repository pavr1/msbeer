package application

import (
	"msbeer.com/adapter"
	"msbeer.com/infra"
	"msbeer.com/models"
)

//BeerApplication Interface
type BeerApplication interface {
	SearchBeers() ([]models.BeerItem, error)
	AddBeers() error
	SearchBeerById(ID int) (models.BeerItem, error)
	BoxBeerPriceById(ID int) (models.BeerBox, error)
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
	return []models.BeerItem{}, nil
}

//AddBeers adds a brand new beer into db
func (a BeerApplicationImpl) AddBeers() error {
	return nil
}

//SearchBeerById searches a beer by ID
func (a BeerApplicationImpl) SearchBeerById(ID int) (models.BeerItem, error) {
	return models.BeerItem{}, nil
}

//BoxBeerPriceById searches for box total value by beer ID
func (a BeerApplicationImpl) BoxBeerPriceById(ID int) (models.BeerBox, error) {
	return models.BeerBox{}, nil
}
