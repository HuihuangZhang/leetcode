package problem_11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	input  []int
	expect int
}

func TestMaxArea(t *testing.T) {
	cases := []TestCase{
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expect: 49,
		},
		{
			input:  []int{1, 1},
			expect: 1,
		},
		{
			input:  []int{1, 8, 6, 2, 5, 4, 8, 25, 7},
			expect: 49,
		},
	}

	for _, c := range cases {
		result := maxArea(c.input)
		assert.Equal(t, c.expect, result)
	}
}
