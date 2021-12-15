package infra

import (
	"msbeer.com/models"
)

type BeerInfra interface {
	SearchBeers() ([]models.BeerItem, error)
	AddBeers() error
	SearchBeerById(ID int) (models.BeerItem, error)
}

type BeerInfraImpl struct {
}

func NewBeerInfraImpl() BeerInfra {
	return BeerInfraImpl{}
}

//SearchBeers searches all list of beers existent in db
func (a BeerInfraImpl) SearchBeers() ([]models.BeerItem, error) {
	return []models.BeerItem{}, nil
}

//AddBeers adds a brand new beer into db
func (a BeerInfraImpl) AddBeers() error {
	return nil
}

//SearchBeerById searches a beer by ID
func (a BeerInfraImpl) SearchBeerById(ID int) (models.BeerItem, error) {
	return models.BeerItem{}, nil
}
