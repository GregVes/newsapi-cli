package utilities

import (
	"testing"

	parser "github.com/gregves/newsapi-golang-client/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestGetStructKeysNames(t *testing.T) {
	tests := []struct {
		input GenericStruct
		want  []string
	}{
		{
			input: &parser.Args{
				QueryAll: "global warming",
				Max:      5,
				Domains:  []string{"bbc.com"},
			},
			// NOTE: order of this array == order of the Arg struct
			want: []string{"QueryAll", "QueryTitle", "Max", "Language", "SortBy", "Domains"},
		},
		{
			input: &parser.Args{
				QueryAll: "global warming",
			},
			// NOTE: order of this array == order of the Arg struct
			want: []string{"QueryAll", "QueryTitle", "Max", "Language", "SortBy", "Domains"},
		},
	}
	for _, tc := range tests {
		got := GetStructKeysNames(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
