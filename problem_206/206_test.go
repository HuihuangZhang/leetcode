package problem206

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want []int
	} {
		{
			name: "case_0",
			input: []int{1,2,3,4,5},
			want: []int{5,4, 3, 2, 1},
		},
		{
			name: "case_1",
			input: []int{1,2},
			want: []int{2,1},
		},
		{
			name: "case_2",
			input: []int{},
			want: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			in := common.FromSlice(tt.input)
			result := reverseList(in.Head)
			
			list := &common.List[int]{Head: result}
			assert.ElementsMatch(t, list.ToSlice(), tt.want)
		})
	}
}