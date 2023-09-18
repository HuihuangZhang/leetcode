package problem94

import (
	"testing"
	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestInorderTraversalIteration(t *testing.T) {
	testCases := []struct {
		name string
		inputTree []int
		wantSlice []int
	} {
		{
			name: "case_0",
			inputTree: []int{1,-101,2,3},
			wantSlice: []int{1,3,2},
		},
		{
			name: "case_1",
			inputTree: []int{},
			wantSlice: []int{},
		},
		{
			name: "case_2",
			inputTree: []int{1},
			wantSlice: []int{1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tmpTree := tree.ImportFromSlice(tt.inputTree, -101)
			result := inorderTraversalIteration(tmpTree)
			assert.ElementsMatch(t, result, tt.wantSlice)
		})
	}
}