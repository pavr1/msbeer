package infra

import (
	"context"
	"database/sql"
	"errors"
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
	db *sql.DB
}

func NewBeerInfraImpl() (BeerInfra, error) {
	connectionString := "server=PAVILLALOBOS;user id=;trusted_connection=true;database=msbeer;app name=msbeer"
	//fmt.Sprintf("Server=localhost;user id=;Database=%s;Trusted_Connection=True;", "msbeer")

	sqlObj, err := sql.Open("mssql", connectionString)
	if err != nil {
		return nil, err
	}

	beerInfra := BeerInfraImpl{
		db: sqlObj,
	}

	return beerInfra, nil
}

//SearchBeers searches all list of beers existent in db
func (a BeerInfraImpl) SearchBeers(ctx context.Context) ([]models.BeerItem, error) {
	result := []models.BeerItem{}
	statement := "SELECT ID, Name, Brewery, Country, Price, Currency FROM beer_item"

	rows, err := a.retrieve(ctx, statement)
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

	err := a.execute(ctx, statement)
	if err != nil {
		return err
	}

	return nil
}

//SearchBeerById searches a beer by ID
func (a BeerInfraImpl) SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error) {
	statement := fmt.Sprintf("SELECT ID, Name, Brewery, Country, Price, Currency FROM beer_item WHERE ID=%d", ID)

	rows, err := a.retrieve(ctx, statement)
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

		return result, nil
	}

	return nil, errors.New("No records found")
}

func (i BeerInfraImpl) execute(ctx context.Context, sqlStatement string) error {
	var err error

	err = i.db.PingContext(ctx)
	if err != nil {
		return err
	}

	query, err := i.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer query.Close()
	newRecord := query.QueryRowContext(ctx)

	var newID int64
	err = newRecord.Scan(&newID)
	if err != nil {
		return err
	}

	return nil
}

func (i BeerInfraImpl) retrieve(ctx context.Context, statement string) (*sql.Rows, error) {
	ctx1 := context.Background()
	err := i.db.PingContext(ctx1)
	if err != nil {
		return nil, err
	}

	data, err := i.db.QueryContext(ctx, statement)
	if err != nil {
		return nil, err
	}

	return data, nil
}
