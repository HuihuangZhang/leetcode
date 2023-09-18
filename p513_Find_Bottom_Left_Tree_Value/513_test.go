package p513findbottomlefttreevalue

import (
	"math"
	"testing"

	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/stretchr/testify/assert"
)

func TestFindBottomLeftValue(t *testing.T) {
	testCases := []struct {
		name string
		input []int
		want int
	} {
		{
			name: "case_0",
			input: []int{2,1,3},
			want: 1,
		},
		{
			name: "case_1",
			input: []int{1,2,3,4,math.MaxInt64,5,6,math.MaxInt64,math.MaxInt64,7},
			want: 7,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := findBottomLeftValue(tree.ImportFromSlice(tt.input, math.MaxInt64))
			assert.Equal(t, tt.want, result)
		})
	}
}