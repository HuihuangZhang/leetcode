package p98validatebinarysearchtree

import (
	"math"
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want bool
	} {
		{
			name: "case_0",
			input: []int{2,1,3},
			want: true,
		},
		{
			name: "case_1",
			input: []int{5,1,4,math.MaxInt64,math.MaxInt64,3,6},
			want: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidBST(tree.ImportFromSlice(tt.input, math.MaxInt64))
			assert.Equal(t, tt.want, result)
		})
	}
}