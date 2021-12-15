package application

import "github.com/stretchr/testify/suite"

type ApplicationTestSuite struct {
	suite.Suite
}

func (s *ApplicationTestSuite) SetupSuite() {
	//instantiate mocks here
}

func (s ApplicationTestSuite) AfterTest(suiteName, testName string) {
	//s.Mock.AssertExpectations(s.T)
	//s.Mock.ExpectedCalls = nil
}

func (s *ApplicationTestSuite) TestHandleSuccess() {
	//arrange

	//act

	//assert
}
