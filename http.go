package main

import (
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

type client struct {
	handlerPath string
	httpClient  *http.Client
}

// New Geter ...
func New(httpClient *http.Client) Geter {
	return &client{
		httpClient: httpClient,
	}
}

// Get ...
func (c *client) Get(url string) (*string, error) {

	client := retryablehttp.NewClient()
	client.HTTPClient = c.httpClient
	req, err := retryablehttp.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseString := string(responseData)
	return &responseString, nil

}
