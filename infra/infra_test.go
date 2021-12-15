package infra

import "github.com/stretchr/testify/suite"

type InfraTestSuite struct {
	suite.Suite
}

func (s *InfraTestSuite) SetupSuite() {
	//instantiate mocks here
}

func (s InfraTestSuite) AfterTest(suiteName, testName string) {
	//s.Mock.AssertExpectations(s.T)
	//s.Mock.ExpectedCalls = nil
}

func (s *InfraTestSuite) TestHandleSuccess() {
	//arrange

	//act

	//assert
}
