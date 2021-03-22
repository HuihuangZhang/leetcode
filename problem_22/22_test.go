package problem_22

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	input  int
	expect []string
}

func TestGenerateParenthesis(t *testing.T) {
	cases := []TestCases{
		{
			input:  3,
			expect: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			input:  4,
			expect: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
		{
			input:  1,
			expect: []string{"()"},
		},
	}

	for _, c := range cases {
		result := generateParenthesis(c.input)
		sort.Strings(result)
		sort.Strings(c.expect)

		assert.Equal(t, c.expect, result)
	}
}
