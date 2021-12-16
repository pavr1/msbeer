package infra

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"msbeer.com/src/models"
)

type InfraUnitTestSuite struct {
	suite.Suite
	BeerInfraImpl BeerInfra
	connMock      DbConnectorMock
}

func TestInfraUnitTestSuite(t *testing.T) {
	suite.Run(t, new(InfraUnitTestSuite))
}

func (s *InfraUnitTestSuite) SetupSuite() {
	s.connMock = DbConnectorMock{}
	beer, _ := NewBeerInfraImpl(&sql.DB{}, s.connMock)

	s.BeerInfraImpl = beer
}

func (s *InfraUnitTestSuite) AfterTest(suiteName, testName string) {
	s.connMock.AssertExpectations(s.T())
	s.connMock.ExpectedCalls = nil
}

type DbConnectorMock struct {
	mock.Mock
}

func (b DbConnectorMock) Execute(ctx context.Context, db *sql.DB, sqlStatement string) error {
	args := b.Called()

	return args.Error(0)
}

func (b DbConnectorMock) Retrieve(ctx context.Context, db *sql.DB, sqlStatement string) ([]models.BeerItem, error) {
	args := b.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	} else {
		return args.Get(0).([]models.BeerItem), args.Error(1)
	}
}

func (s *InfraUnitTestSuite) TestNewBeerInfraImplSuccess() {
	//arrange
	s.connMock = DbConnectorMock{}
	db := &sql.DB{}
	beerInfra, err := NewBeerInfraImpl(db, s.connMock)
	//act

	//assert
	s.Assert().Nil(err)
	s.Assert().IsType(BeerInfraImpl{}, beerInfra)
}

func (s *InfraUnitTestSuite) Test_SearchBeers_Retrieve_Success() {
	ctx := context.Background()

	beerItems := []models.BeerItem{
		{
			ID:       1,
			Name:     "Pilsen",
			Brewery:  "test",
			Country:  "Costa Rica",
			Price:    2.5,
			Currency: "USD",
		},
	}

	s.connMock.On("Retrieve").Return(beerItems, nil)
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	list, err := s.BeerInfraImpl.SearchBeers(ctx)
	s.Assert().Nil(err)
	s.Assert().Equal(len(list), 1)
	s.Assert().Contains(list[0].Name, "Pilsen")
}

func (s *InfraUnitTestSuite) Test_SearchBeers_Retrieve_Failed() {
	ctx := context.Background()

	s.connMock.On("Retrieve").Return(nil, errors.New("fake error"))
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem, err := s.BeerInfraImpl.SearchBeers(ctx)
	s.Assert().NotNil(err)
	s.Assert().Contains(err.Error(), "fake error")
	s.Assert().Nil(beerItem)
}

func (s *InfraUnitTestSuite) Test_AddBeers_Execute_Success() {
	ctx := context.Background()

	s.connMock.On("Execute").Return(nil)
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem := models.BeerItem{
		ID:       1,
		Name:     "Pilsen",
		Brewery:  "test",
		Country:  "Costa Rica",
		Price:    2.5,
		Currency: "USD",
	}

	err = s.BeerInfraImpl.AddBeers(ctx, beerItem)
	s.Assert().Nil(err)
}

func (s *InfraUnitTestSuite) Test_AddBeers_Execute_Failed() {
	ctx := context.Background()

	s.connMock.On("Execute").Return(errors.New("fake error"))
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem := models.BeerItem{
		ID:       1,
		Name:     "Pilsen",
		Brewery:  "test",
		Country:  "Costa Rica",
		Price:    2.5,
		Currency: "USD",
	}

	err = s.BeerInfraImpl.AddBeers(ctx, beerItem)
	s.Assert().NotNil(err)
	s.Assert().Contains(err.Error(), "fake error")
}

func (s *InfraUnitTestSuite) Test_SearchBeerById_Retrieve_Success() {
	ctx := context.Background()

	beerItems := []models.BeerItem{
		{
			ID:       1,
			Name:     "Pilsen",
			Brewery:  "test",
			Country:  "Costa Rica",
			Price:    2.5,
			Currency: "USD",
		},
	}

	s.connMock.On("Retrieve").Return(beerItems, nil)
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem, err := s.BeerInfraImpl.SearchBeerById(ctx, 1)
	s.Assert().Nil(err)
	s.Assert().NotNil(beerItem)
	s.Assert().Contains(beerItem.Name, "Pilsen")
}

func (s *InfraUnitTestSuite) Test_SearchBeerById_Retrieve_Empty_Result() {
	ctx := context.Background()

	beerItems := []models.BeerItem{}

	s.connMock.On("Retrieve").Return(beerItems, nil)
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem, err := s.BeerInfraImpl.SearchBeerById(ctx, 1)
	s.Assert().NotNil(err)
	s.Assert().Contains(err.Error(), "no items found")
	s.Assert().Nil(beerItem)
}

func (s *InfraUnitTestSuite) Test_SearchBeerById_Retrieve_Failed() {
	ctx := context.Background()

	s.connMock.On("Retrieve").Return(nil, errors.New("fake error"))
	beerInfra, err := NewBeerInfraImpl(&sql.DB{}, s.connMock)
	s.Assert().Nil(err)
	s.BeerInfraImpl = beerInfra

	beerItem, err := s.BeerInfraImpl.SearchBeerById(ctx, 1)
	s.Assert().NotNil(err)
	s.Assert().Contains(err.Error(), "fake error")
	s.Assert().Nil(beerItem)
}
