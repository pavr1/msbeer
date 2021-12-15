package handlers

import "github.com/stretchr/testify/suite"

type HandlerTestSuite struct {
	suite.Suite
}

func (s *HandlerTestSuite) SetupSuite() {
	//instantiate mocks here
}

func (s HandlerTestSuite) AfterTest(suiteName, testName string) {
	//s.Mock.AssertExpectations(s.T)
	//s.Mock.ExpectedCalls = nil
}

func (s *HandlerTestSuite) TestHandleSuccess() {
	//arrange

	//act

	//assert
}
