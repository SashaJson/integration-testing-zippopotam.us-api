package tests

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/suite"
)

type ZippopotamUsTestSuite struct {
	suite.Suite
	ApiClient *resty.Client
}