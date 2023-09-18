package problem257

import (
	"sort"
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreePaths(t *testing.T) {
	testCases := []struct {
		name string
		inputTree []int
		want []string
	} {
		{
			name: "case_0",
			inputTree: []int{1,2,3,-101,5},
			want: []string{"1->2->5","1->3"},
		},
		{
			name: "case_1",
			inputTree: []int{1},
			want: []string{"1"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := binaryTreePaths(tree.ImportFromSlice(tt.inputTree, -101))
			sort.Strings(result)
			sort.Strings(tt.want)
			assert.ElementsMatch(t, result, tt.want)
		})
	}
}
