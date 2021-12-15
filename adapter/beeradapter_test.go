package adapter

import "github.com/stretchr/testify/suite"

type AdapterTestSuite struct {
	suite.Suite
}

func (s *AdapterTestSuite) SetupSuite() {
	//instantiate mocks here
}

func (s AdapterTestSuite) AfterTest(suiteName, testName string) {
	//s.Mock.AssertExpectations(s.T)
	//s.Mock.ExpectedCalls = nil
}

func (s *AdapterTestSuite) TestHandleSuccess() {
	//arrange

	//act

	//assert
}
