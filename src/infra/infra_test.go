package infra

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type InfraUnitTestSuite struct {
	suite.Suite
	BeerInfraImpl BeerInfra
	connMock      DbConnectorMock
}

type DbConnectorMock struct {
	mock.Mock
}

func (b DbConnectorMock) Execute(ctx context.Context, db *sql.DB, sqlStatement string) error {
	args := b.Called()

	return args.Error(0)
}

func (b DbConnectorMock) Retrieve(ctx context.Context, db *sql.DB, statement string) (*sql.Rows, error) {
	args := b.Called()

	return args.Get(0).(*sql.Rows), args.Error(1)
}

func (s *InfraUnitTestSuite) SetupSuite() {
	s.connMock = DbConnectorMock{}
	beer, _ := NewBeerInfraImpl(&sql.DB{}, s.connMock)

	s.BeerInfraImpl = beer
}

func (s *InfraUnitTestSuite) SetupTest() {

}

func (s *InfraUnitTestSuite) AfterTest(suiteName, testName string) {
	s.connMock.AssertExpectations(s.T())
	s.connMock.ExpectedCalls = nil
}

func (s *InfraUnitTestSuite) TestInfraUnitTestSuite(t *testing.T) {
	suite.Run(t, &InfraUnitTestSuite{})
}

func (s *InfraUnitTestSuite) TestNewBeerInfraImplSuccess() {
	//arrange
	s.connMock = DbConnectorMock{}
	db := &sql.DB{}
	beerInfra, err := NewBeerInfraImpl(db, s.connMock)
	//act

	//assert
	s.Assert().Nil(err)
	s.Assert().IsType(reflect.TypeOf(BeerInfraImpl{}), beerInfra)
}

func (s *InfraUnitTestSuite) TestSearchBeersSuccess() {
	ctx := context.Background()
	rows := []*sql.Rows{}
	rows = append(rows, &sql.Rows{})
	rows = append(rows, &sql.Rows{})
	rows = append(rows, &sql.Rows{})

	s.connMock.On("Retrieve").Return(rows, nil)

	beerItem, err := s.BeerInfraImpl.SearchBeers(ctx)

	s.Assert().Nil(err)
	s.Assert().Equal(3, len(beerItem))
}
