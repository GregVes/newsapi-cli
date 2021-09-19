package query

import (
	"testing"

	"github.com/gregves/newsapi-golang-client/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestConvertArgsToStringQuery(t *testing.T) {
	tests := []struct {
		input parser.Args
		want  string
	}{
		{
			input: parser.Args{
				QueryTitle: "global warming",
				Language:   "it",
				Domains:    []string{"lemonde.fr", "bbc.com"},
			},
			want: "qInTitle=global%20warming&language=it&domains=lemonde.fr,bbc.com&page=1",
		},
		{
			input: parser.Args{
				QueryTitle: "amazonas wald",
				Language:   "de",
			},
			want: "qInTitle=amazonas%20wald&language=de&page=1",
		},
		{
			input: parser.Args{
				QueryTitle: "justin bieber",
				SortBy:     "popularity",
			},
			want: "qInTitle=justin%20bieber&sortBy=popularity&page=1",
		},
		{
			input: parser.Args{
				QueryAll: "mighty vaporizer",
				Max:      5,
			},
			want: "q=mighty%20vaporizer&pageSize=5&page=1",
		},
	}
	for _, tc := range tests {
		queryMapper := New(&tc.input)
		queryMapper.ToQuery()
		got := queryMapper.QuerySring
		assert.Equal(t, tc.want, got)
	}
}
