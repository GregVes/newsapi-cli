package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gregves/newsapi-golang-client/pkg/article"
	"github.com/stretchr/testify/assert"
)

const ISO8601_LAYOUT = "2006-01-02T15:04:05Z"

type (
	mockAPIClient struct {
		Client     http.Client
		ApiBaseUrl string
		Query      string
		ApiKey     string
	}
)

func NewMockAPIClient() mockAPIClient {
	return mockAPIClient{
		Client:     http.Client{},
		ApiBaseUrl: "url",
		Query:      "query",
		ApiKey:     "api-key",
	}
}

// will be called instead of client.go:Search()
func (client *mockAPIClient) Search(query string) (*http.Response, error) {
	reponseMockFile, _ := os.ReadFile("./data/response_mock.json")
	reader := ioutil.NopCloser(bytes.NewReader([]byte(reponseMockFile)))
	return &http.Response{
		Body: reader,
	}, nil
}

func TestMockingSearchingArticles(t *testing.T) {
	time, err := time.Parse(ISO8601_LAYOUT, "2021-08-23T13:49:45Z")
	if err != nil {
		log.Println(err)
	}
	want := NewsAPIResponse{
		Status:       "200",
		TotalResults: 10,
		Articles: []article.Article{
			article.Article{
				Source: map[string]string{
					"id":   "0000",
					"name": "super name",
				},
				Author:      "greg",
				Title:       "super bomb",
				Description: "a bomb exploded",
				Url:         "https://super-news.com",
				Content:     "this is the content",
				PublishedAt: time,
			},
		},
	}
	mockClient := NewMockAPIClient()
	service := NewAPIService(&mockClient)

	got, err := service.GetArticles("q=global%20warming&page=1")

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, got, want.Articles)
}
