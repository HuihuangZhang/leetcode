package problem203

import (
	"testing"
	"github.com/HuihuangZhang/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		name string
		inputList []int
		inputVal int
		want []int
	} {
		{
			name: "case_0",
			inputList: []int{1,2,6,3,4,5,6},
			inputVal: 6,
			want: []int{1,2,3,4,5},
		},
		{
			name: "case_1",
			inputList: []int{},
			inputVal: 1,
			want: []int{},
		},
		{
			name: "case_2",
			inputList: []int{7,7,7,7},
			inputVal: 7,
			want: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElements(common.FromSlice(tt.inputList).Head, tt.inputVal)
			resultList := &common.List[int]{
				Head: result,
				Size: 0,
			}
			assert.ElementsMatch(t, resultList.ToSlice(), tt.want)
		})
	}
}
