package infra

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"msbeer.com/src/models"
)

type BeerInfra interface {
	SearchBeers(ctx context.Context) ([]models.BeerItem, error)
	AddBeers(ctx context.Context, beer models.BeerItem) error
	SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error)
}

type BeerInfraImpl struct {
	db   *sql.DB
	conn DbConnector
}

func NewBeerInfraImpl(db *sql.DB, conn DbConnector) (BeerInfra, error) {
	beerInfra := BeerInfraImpl{
		db: db,
	}

	return beerInfra, nil
}

//SearchBeers searches all list of beers existent in db
func (a BeerInfraImpl) SearchBeers(ctx context.Context) ([]models.BeerItem, error) {
	result := []models.BeerItem{}
	statement := "SELECT ID, Name, Brewery, Country, Price, Currency FROM beer_item"

	rows, err := a.conn.Retrieve(ctx, a.db, statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var name, brewery, country, currency string
		var id int
		var price float64

		err = rows.Scan(&id, &name, &brewery, &country, &price, &currency)
		if err != nil {
			return nil, err
		}

		result = append(result, models.BeerItem{
			ID:       id,
			Name:     name,
			Brewery:  brewery,
			Country:  country,
			Price:    price,
			Currency: currency,
		})
	}

	return result, nil
}

//AddBeers adds a brand new beer into db
func (a BeerInfraImpl) AddBeers(ctx context.Context, beer models.BeerItem) error {
	statement := fmt.Sprintf("INSERT INTO beer_item(ID, Name, Brewery, Country, Price, Currency) VALUES (%d, '%s', '%s', '%s', %v, '%s')",
		beer.ID,
		beer.Name,
		beer.Brewery,
		beer.Country,
		beer.Price,
		beer.Currency)

	err := a.conn.Execute(ctx, a.db, statement)
	if err != nil {
		return err
	}

	return nil
}

//SearchBeerById searches a beer by ID
func (a BeerInfraImpl) SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error) {
	statement := fmt.Sprintf("SELECT ID, Name, Brewery, Country, Price, Currency FROM beer_item WHERE ID=%d", ID)

	rows, err := a.conn.Retrieve(ctx, a.db, statement)
	if err != nil {
		return nil, err
	}

	var result *models.BeerItem
	for rows.Next() {
		var name, brewery, country, currency string
		var id int
		var price float64

		err = rows.Scan(&id, &name, &brewery, &country, &price, &currency)
		if err != nil {
			return nil, err
		}

		result = &models.BeerItem{
			ID:       id,
			Name:     name,
			Brewery:  brewery,
			Country:  country,
			Price:    price,
			Currency: currency,
		}

		break
	}

	return result, nil
}
