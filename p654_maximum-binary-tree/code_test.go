package p654maximumbinarytree

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestConstructMaximumBinaryTree(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want []int
	} {
		{
			name: "case_0",
			input: []int{3,2,1},
			want: []int{3,-1,2,-1,1},
		},
		{
			name: "case_1",
			input: []int{3,2,1,6,0,5},
			want: []int{6,3,5,-1,2,0,-1,-1,1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := constructMaximumBinaryTree(tt.input)
			assert.ElementsMatch(t, tree.ExportToSliceByLevelTraversal(result, -1), tt.want)
		})
	}

}