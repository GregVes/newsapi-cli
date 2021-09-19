package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingCommandLineArguments(t *testing.T) {

	tests := []struct {
		input []string
		want  *Args
	}{
		{
			input: []string{"main", "-q", "global warming", "-d", "lemonde.fr"},
			want: &Args{
				QueryAll: "global warming",
				Domains:  []string{"lemonde.fr"},
			},
		},
		{
			input: []string{"main", "-t", "super title"},
			want: &Args{
				QueryTitle: "super title",
				Domains:    []string{},
			},
		},
		{
			input: []string{"main", "-q", "sebastiao salgado", "-l", "it"},
			want: &Args{
				QueryAll: "sebastiao salgado",
				Domains:  []string{},
				Language: "it",
			},
		},
		{
			input: []string{"main", "-t", "mario bros", "-l", "fr", "-s", "popularity", "-d", "bbc.com", "-d", "theguardian.com"},
			want: &Args{
				QueryTitle: "mario bros",
				Domains:    []string{"bbc.com", "theguardian.com"},
				Language:   "fr",
				SortBy:     "popularity",
			},
		},
	}
	for _, tc := range tests {
		parser := New()
		parser.Parse(tc.input)
		got := parser.Args
		assert.Equal(t, tc.want, got)
	}
}
