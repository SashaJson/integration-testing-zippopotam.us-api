package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_First(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))

}
