package problem_19

import (
	"testing"

	"github.com/HuihuangZhang/leetcode/common"
	"github.com/stretchr/testify/assert"
)

type input struct {
	list []int
	n    int
}

type testCases struct {
	in     input
	expect []int
}

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []testCases{
		{
			in: input{
				list: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			expect: []int{1, 2, 3, 5},
		},
		{
			in: input{
				list: []int{1},
				n:    1,
			},
			expect: []int{},
		},
		{
			in: input{
				list: []int{1, 2},
				n:    1,
			},
			expect: []int{1},
		},
		{
			in: input{
				list: []int{1, 2},
				n:    2,
			},
			expect: []int{2},
		},
	}

	for _, c := range cases {
		l := common.FromSlice(c.in.list)
		head := removeNthFromEnd(l.Head, c.in.n)
		resultL := common.List[int]{Head: head}
		assert.Equal(t, c.expect, resultL.ToSlice())
	}
}
