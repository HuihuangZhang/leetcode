package p538convertbsttogreatertree

import (
	"math"
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

const null = math.MaxInt64

func TestConvertBST(t *testing.T) {
	testCases := []struct{
		name string
		input []int
		want []int
	} {
		{
			name: "case_0",
			input: []int{4,1,6,0,2,5,7,null,null,null,3,null,null,null,8},
			want: []int{30,36,21,36,35,26,15,null,33,null,8},
		},
		{
			name: "case_1",
			input: []int{0,null,1},
			want: []int{1,null,1},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := convertBST(tree.ImportFromSlice(tt.input, null))
			assert.ElementsMatch(t, tree.ExportToSliceByLevelTraversal(result, null), tt.want)
		})
	}
}
