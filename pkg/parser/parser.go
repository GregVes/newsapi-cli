package parser

import (
	"log"

	"github.com/akamensky/argparse"
)

type (
	Args struct {
		QueryAll   string
		QueryTitle string
		Max        int
		Language   string
		SortBy     string
		Domains    []string
	}
	Parser struct {
		Args   *Args
		Parser argparse.Parser
	}
)

var (
	cliHelp        = "search for news and display them in your browser"
	queryAllHelp   = "String to search in news titles and bodies"
	queryTitleHelp = "String to search in news titles"
	languageHelp   = "Language of news to search for"
	sortByHelp     = "Choose from: relevancy, popularity or publishedAt (default)"
	domainsHelp    = "List of domains to search for news in"
	maxHelp        = "Maximum articles to retrieve"

	sortBy = []string{"relevancy", "popularity", "publishedAt"}
)

func New() *Parser {
	return &Parser{
		Args:   &Args{},
		Parser: *argparse.NewParser("news", cliHelp),
	}
}

func (p *Parser) Parse(args []string) error {
	// create args template

	queryAll := p.Parser.String("q", "query", &argparse.Options{Required: false, Help: queryAllHelp})
	queryTitle := p.Parser.String("t", "title", &argparse.Options{Required: false, Help: queryTitleHelp})
	max := p.Parser.Int("m", "max", &argparse.Options{Required: false, Help: maxHelp})
	language := p.Parser.String("l", "language", &argparse.Options{Required: false, Help: languageHelp})
	domains := p.Parser.StringList("d", "domains", &argparse.Options{Required: false, Help: domainsHelp})
	sortBy := p.Parser.Selector("s", "sortby", sortBy, &argparse.Options{Required: false, Help: sortByHelp})

	// parse
	err := p.Parser.Parse(args)
	if err != nil {
		log.Println(p.Parser.Usage(err))
		return err
	}

	p.Args = &Args{
		QueryAll:   *queryAll,
		QueryTitle: *queryTitle,
		Max:        *max,
		Language:   *language,
		Domains:    *domains,
		SortBy:     *sortBy,
	}

	return nil
}
