package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gregves/newsapi-golang-client/pkg/article"
	client "github.com/gregves/newsapi-golang-client/pkg/client"
)

type (
	APIService struct {
		Client client.APIClient
	}
	NewsAPIResponse struct {
		Status       string
		TotalResults int
		Articles     []article.Article
	}
)

func NewAPIService(apiClient client.APIClient) *APIService {
	return &APIService{apiClient}
}

func (service *APIService) GetArticles(query string) ([]article.Article, error) {
	// Search() is mocked in unit tests
	response, err := service.Client.Search(query)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	// convert ReadCloser into byte[] array
	bodyByte, e := ioutil.ReadAll(response.Body)

	if e != nil {
		log.Println(e)
		return nil, e
	}

	// unmarshall byte array and convert into []Article
	var newsAPIResponse NewsAPIResponse
	e = json.Unmarshal([]byte(bodyByte), &newsAPIResponse)

	if e != nil {
		fmt.Println(e)
		return nil, e
	}

	return newsAPIResponse.Articles, nil
}
