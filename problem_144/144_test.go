package problem144

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []struct {
		name string
		inputTree []int
		wantSlice []int
	} {
		{
			name: "case_0",
			inputTree: []int{1,-101,2,3},
			wantSlice: []int{1,2,3},
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
		{
			name: "case_3",
			inputTree: []int{1, 2},
			wantSlice: []int{1, 2},
		},
		{
			name: "case_4",
			inputTree: []int{1, -101, 2},
			wantSlice: []int{1, 2},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tmpTree := tree.ImportFromSlice(tt.inputTree, -101)
			result := preorderTraversal(tmpTree)
			assert.ElementsMatch(t, result, tt.wantSlice)

			result = preorderTraversalIteration(tmpTree)
			assert.ElementsMatch(t, result, tt.wantSlice)
		})

		t.Run(tt.name + "_iteration", func(t *testing.T) {
			tmpTree := tree.ImportFromSlice(tt.inputTree, -101)
			result := preorderTraversalIteration(tmpTree)
			assert.ElementsMatch(t, result, tt.wantSlice)
		})
	}
}