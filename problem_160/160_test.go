package problem160

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	testCases := []struct{
		name string
		inputA []int
		skipA int
		inputB []int
		skipB int
		want int
	} {
		{
			name: "case_0",
			inputA: []int{4,1,8,4,5},
			inputB: []int{5,0,1,8,4,5},
			skipA: 2,
			skipB: 3,
			want: 8,
		},
		{
			name: "case_1",
			inputA: []int{0,9,1,2,4},
			inputB: []int{3,2,4},
			skipA: 3,
			skipB: 1,
			want: 2,
		},
		{
			name: "case_2",
			inputA: []int{2,6,4},
			inputB: []int{1,5},
			skipA: 3,
			skipB: 2,
			want: 0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			listA, listB := helper(tt.inputA, tt.inputB, tt.skipA, tt.skipB)
			commonHead := getIntersectionNode(listA, listB)

			if tt.want == 0 {
				assert.Nil(t, commonHead)
			} else {
				assert.Equal(t, tt.want, commonHead.Val)
			}
		})
	}
}

func helper(inputA, inputB []int, skipA, skipB int) (*common.ListNode[int], *common.ListNode[int]) {
	lA := common.FromSlice(inputA[:skipA]).Head
	lB := common.FromSlice(inputB[:skipB]).Head

	common := common.FromSlice(inputA[skipA:]).Head

	lA.Next = common
	lB.Next = common

	return lA, lB
}