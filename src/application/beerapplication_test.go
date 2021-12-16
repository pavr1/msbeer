package application

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"msbeer.com/src/models"
)

type BeerApplicationUnitTestSuite struct {
	suite.Suite
	BeerApplicationImpl BeerApplicationImpl
	Adapter             BeerAdapterMock
	Infrastructure      BeerInfraMock
}

type BeerAdapterMock struct {
	mock.Mock
}

func (b BeerAdapterMock) GetCurrencyInfo() (*models.Currency, error) {
	args := b.Called()

	return args.Get(0).(*models.Currency), args.Error(1)
}

type BeerInfraMock struct {
	mock.Mock
}

func (b BeerInfraMock) SearchBeers(ctx context.Context) ([]models.BeerItem, error) {
	args := b.Called()

	return args.Get(0).([]models.BeerItem), args.Error(1)
}
func (b BeerInfraMock) AddBeers(ctx context.Context, beer models.BeerItem) error {
	args := b.Called()

	return args.Error(0)
}
func (b BeerInfraMock) SearchBeerById(ctx context.Context, ID int) (*models.BeerItem, error) {
	args := b.Called()

	return args.Get(0).(*models.BeerItem), args.Error(1)
}

func TestBeerApplicationUnitTestSuite(t *testing.T) {
	suite.Run(t, new(BeerApplicationUnitTestSuite))
}

func (s *BeerApplicationUnitTestSuite) SetupSuite() {
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

	s.BeerApplicationImpl = BeerApplicationImpl{s.Adapter, s.Infrastructure}
}

func (s *BeerApplicationUnitTestSuite) AfterTest(suiteName, testName string) {
	s.Adapter.AssertExpectations(s.T())
	s.Adapter.ExpectedCalls = nil

	s.Infrastructure.AssertExpectations(s.T())
	s.Infrastructure.ExpectedCalls = nil
}

func (s *BeerApplicationUnitTestSuite) Test_NewApplication_Success() {
	//arrange
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

	appl := NewApplication(s.Adapter, s.Infrastructure)
	//act

	//assert
	s.Assert().IsType(BeerApplicationImpl{}, appl)
}

func (s *BeerApplicationUnitTestSuite) Test_SearchBeers_Success() {
	//arrange
	ctx := context.Background()
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

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

	//act
	s.Infrastructure.On("SearchBeers").Return(beerItems, nil)
	appl := NewApplication(s.Adapter, s.Infrastructure)

	list, err := appl.SearchBeers(ctx)
	//assert
	s.Assert().Nil(err)
	s.Assert().Equal(1, len(list))
}

func (s *BeerApplicationUnitTestSuite) Test_AddBeers_Success() {
	//arrange
	ctx := context.Background()
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

	beerItem := models.BeerItem{
		ID:       1,
		Name:     "Pilsen",
		Brewery:  "test",
		Country:  "Costa Rica",
		Price:    2.5,
		Currency: "USD",
	}

	//act
	s.Infrastructure.On("AddBeers").Return(nil)
	appl := NewApplication(s.Adapter, s.Infrastructure)

	err := appl.AddBeers(ctx, beerItem)
	//assert
	s.Assert().Nil(err)
}

func (s *BeerApplicationUnitTestSuite) Test_SearchBeerById_Success() {
	//arrange
	ctx := context.Background()
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

	beerItem := models.BeerItem{
		ID:       1,
		Name:     "Pilsen",
		Brewery:  "test",
		Country:  "Costa Rica",
		Price:    2.5,
		Currency: "USD",
	}

	//act
	s.Infrastructure.On("SearchBeerById").Return(&beerItem, nil)
	appl := NewApplication(s.Adapter, s.Infrastructure)

	item, err := appl.SearchBeerById(ctx, 1)
	//assert
	s.Assert().Nil(err)
	s.Assert().NotNil(item)
	s.Assert().Equal("Pilsen", item.Name)
}

func (s *BeerApplicationUnitTestSuite) Test_BoxBeerPriceById_Success() {
	//arrange
	ctx := context.Background()
	s.Adapter = BeerAdapterMock{}
	s.Infrastructure = BeerInfraMock{}

	currencyInfo := models.Currency{
		Success:   true,
		Terms:     "",
		Privacy:   "",
		Timestamp: time.Time{}.Unix(),
		Source:    "",
		Quote: models.Quote{
			USDAED: 1,
			USDAFN: 2,
			USDALL: 3,
			USDAMD: 4,
			USDANG: 5,
			USDAOA: 6,
			USDARS: 7,
			USDAUD: 8,
			USDAWG: 9,
			USDAZN: 10,
			USDBAM: 11,
		},
	}
	beerItem := models.BeerItem{
		ID:       1,
		Name:     "Pilsen",
		Brewery:  "test",
		Country:  "Costa Rica",
		Price:    2.5,
		Currency: "USD",
	}

	//act
	s.Adapter.On("GetCurrencyInfo").Return(&currencyInfo, nil)
	s.Infrastructure.On("SearchBeerById").Return(&beerItem, nil)
	appl := NewApplication(s.Adapter, s.Infrastructure)

	item, err := appl.BoxBeerPriceById(ctx, 1, 5, "USDAUD")
	//assert
	s.Assert().Nil(err)
	s.Assert().NotNil(item)
	s.Assert().Equal(item.PriceTotal, float64(100))
}
