package problem_24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	input, expect []int
}

func TestSwapPairs(t *testing.T) {
	cases := []TestCases{
		{
			input:  []int{1, 2, 3, 4},
			expect: []int{2, 1, 4, 3},
		},
		{
			input:  []int{},
			expect: []int{},
		},
		{
			input:  []int{1},
			expect: []int{1},
		},
	}
	for _, c := range cases {
		list := arr2List(c.input)
		result := list2Arr(swapPairs(list))
		assert.Equal(t, c.expect, result, "should be the same")
	}
}
