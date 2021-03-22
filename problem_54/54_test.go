package problem_54

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	input  [][]int
	expect []int
}

func TestSpiralOrder(t *testing.T) {
	cases := []TestCases{{
		input:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
		expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	}

	for _, c := range cases {
		result := spiralOrder(c.input)
		assert.Equal(t, c.expect, result)
	}
}
