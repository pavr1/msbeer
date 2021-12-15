package infra

import (
	// "database/sql"

	"msbeer.com/src/models"
)

type BeerInfra interface {
	SearchBeers() ([]models.BeerItem, error)
	AddBeers(models.BeerItem) error
	SearchBeerById(ID int) (models.BeerItem, error)
}

type BeerInfraImpl struct {
	//db *sql.DB
}

func NewBeerInfraImpl() (BeerInfra, error) {
	// database, err := sql.Open("sqlite3", "./bogo.db")
	// if err != nil {
	// 	return nil, err
	// }

	// fmt.Println(database)
	// return BeerInfraImpl{
	// 	db: database,
	// }, nil
	return nil, nil
}

//SearchBeers searches all list of beers existent in db
func (a BeerInfraImpl) SearchBeers() ([]models.BeerItem, error) {
	return []models.BeerItem{}, nil
}

//AddBeers adds a brand new beer into db
func (a BeerInfraImpl) AddBeers(beer models.BeerItem) error {
	return nil
}

//SearchBeerById searches a beer by ID
func (a BeerInfraImpl) SearchBeerById(ID int) (models.BeerItem, error) {
	return models.BeerItem{}, nil
}
