package problem404

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestSumOfLeftLeaves(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want int
	} {
		{
			name: "case_0",
			input: []int{3,9,20, -101, -101,15,7},
			want: 24,
		},
		{
			name: "case_1",
			input: []int{1},
			want: 0,
		},
		{
			name: "case_2",
			input: []int{1, 2},
			want: 2,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := sumOfLeftLeaves(tree.ImportFromSlice(tt.input, -101))
			assert.Equal(t, tt.want, result)
		})
	}
}