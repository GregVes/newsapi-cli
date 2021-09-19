package main

import (
	"log"
	"os"
	"sort"

	"github.com/gregves/newsapi-golang-client/pkg/client"
	"github.com/gregves/newsapi-golang-client/pkg/parser"
	pprint "github.com/gregves/newsapi-golang-client/pkg/prettyprint"
	"github.com/gregves/newsapi-golang-client/pkg/query"
	api "github.com/gregves/newsapi-golang-client/pkg/services"
)

var baseUrl = "https://newsapi.org/v2/everything?"

func main() {
	start(os.Args)
}

func start(args []string) {
	apiKey := os.Getenv("NEWS_API_KEY")
	if apiKey == "" {
		log.Println("Missing API key. Get one - https://newsapi.org/register - and export NEWS_API_KEY in your ~/.bashrc or ~/.zshrc")
		os.Exit(1)
	}

	//log.SetFlags(log.LstdFlags | log.Lshortfile)

	p := parser.New()
	err := p.Parse(args)

	if err != nil {
		os.Exit(1)
	}

	queryMapper := query.New(p.Args)
	queryMapper.ToQuery()

	client := client.New(baseUrl, queryMapper.QuerySring, apiKey)
	apiService := api.NewAPIService(client)

	articles, err := apiService.GetArticles(queryMapper.QuerySring)

	if err != nil {
		log.Fatal(err)
	}

	if len(articles) == 0 {
		log.Println("No article found")
		os.Exit(0)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishedAt.Before(articles[j].PublishedAt)
	})

	pprint.PrettyPrint(articles)

}
