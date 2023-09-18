package problem128

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{},
			expect: 0,
		},
		{
			input:  []int{1},
			expect: 1,
		},
		{
			input:  []int{100, 1, 2, 99, 3},
			expect: 3,
		},
		{
			input:  []int{98, 1, 2, 99, 100, 3, 5, 4},
			expect: 5,
		},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			result := longestConsecutive(tt.input)
			assert.Equal(t, tt.expect, result)
		})
	}
}
