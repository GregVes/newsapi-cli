package query

import (
	"fmt"
	"strings"

	"github.com/gregves/newsapi-golang-client/pkg/parser"
)

type (
	QueryMapper struct {
		Args       parser.Args
		QuerySring string
	}
)

var (
	arg2ParamMap = map[string]map[string]interface{}{
		"QueryAll": {
			"key":   0,
			"param": "q=",
			"empty": "",
		},
		"QueryTitle": {
			"key":   1,
			"param": "qInTitle=",
			"empty": "",
		},
		"Max": {
			"key":   2,
			"param": "pageSize=",
			"empty": 0,
		},
		"Language": {
			"key":   3,
			"param": "language=",
			"empty": "",
		},
		"SortBy": {
			"key":   4,
			"param": "sortBy=",
			"empty": "",
		},
	}
)

func New(args *parser.Args) *QueryMapper {
	return &QueryMapper{
		Args: *args,
	}
}

//I cannot send argValue in EmptyField() as I need to know the struct field name
// which is not possible with the iteration we use
func (queryMapper *QueryMapper) ToQuery() {
	var query string
	if queryMapper.Args.QueryAll != "" {
		formattedQueryAll := strings.ReplaceAll(queryMapper.Args.QueryAll, " ", "%20")
		query = query + fmt.Sprintf("q=%s&", formattedQueryAll)
	}

	if queryMapper.Args.QueryTitle != "" {
		formattedQueryTitle := strings.ReplaceAll(queryMapper.Args.QueryTitle, " ", "%20")
		query = query + fmt.Sprintf("qInTitle=%s&", formattedQueryTitle)
	}

	if queryMapper.Args.Max != 0 {
		query = query + fmt.Sprintf("pageSize=%d&", queryMapper.Args.Max)
	}

	if queryMapper.Args.Language != "" {
		query = query + fmt.Sprintf("language=%s&", queryMapper.Args.Language)
	}

	if len(queryMapper.Args.Domains) != 0 {
		query = query + "domains="
		for i, domain := range queryMapper.Args.Domains {
			param := fmt.Sprintf("%s", domain)
			query = query + param
			if i < len(queryMapper.Args.Domains)-1 {
				query = query + ","
			}
		}
		query = query + "&"
	}

	if queryMapper.Args.SortBy != "" {
		query = query + fmt.Sprintf("sortBy=%s&", queryMapper.Args.SortBy)
	}

	// by default, we display all articles in one page
	query = query + "page=1"

	queryMapper.QuerySring = query
}
