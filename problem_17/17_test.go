package problem_17

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	input  string
	expect []string
}

func TestLetterCombinations(t *testing.T) {
	cases := []TestCases{
		{
			input:  "23",
			expect: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input:  "",
			expect: []string{},
		},
		{
			input:  "2",
			expect: []string{"a", "b", "c"},
		},
	}

	for _, c := range cases {
		result := letterCombinations(c.input)
		sort.Strings(c.expect)
		sort.Strings(result)
		assert.Equal(t, c.expect, result)
	}
}
