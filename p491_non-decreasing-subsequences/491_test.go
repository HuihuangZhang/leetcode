package p491nondecreasingsubsequences

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSubsequences(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want [][]int
	} {
		{
			name: "case_0",
			input: []int{4,6,7,7},
			want: [][]int{{4,6},{4,6,7},{4,6,7,7},{4,7},{4,7,7},{6,7},{6,7,7},{7,7}},
		},
		{
			name: "case_1",
			input: []int{4,4,3,2,1},
			want: [][]int{{4,4}},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubsequences(tt.input)
			assert.Equal(t, len(result), len(tt.want))

			for _, item := range result {
				assert.Contains(t, tt.want, item)
			}
		})
	}
}