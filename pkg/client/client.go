package client

import (
	"net/http"
)

type (
	APIClient interface {
		Search(query string) (*http.Response, error)
	}

	NewsAPIClient struct {
		Client     http.Client
		ApiBaseUrl string
		Query      string
		ApiKey     string
	}
)

// we use the NewsAPIClient interface as a return value type to mock http requests in the unit tests
func New(apiBaseUrl string, query string, apiKey string) *NewsAPIClient {
	return &NewsAPIClient{
		Client:     http.Client{},
		ApiBaseUrl: apiBaseUrl,
		Query:      query,
		ApiKey:     apiKey,
	}
}

// we use the NewsAPIClient interface as a receiver type to mock http requests in the unit tests
func (client *NewsAPIClient) Search(query string) (*http.Response, error) {
	endpoint := client.ApiBaseUrl + client.Query + "&apiKey=" + client.ApiKey

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
