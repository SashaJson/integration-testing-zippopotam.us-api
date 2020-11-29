package main

import "github.com/go-resty/resty/v2"

func (suite *ZippopotamUsTestSuite) SetupTest() {
	suite.ApiClient = resty.New()
}