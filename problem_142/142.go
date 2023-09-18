package problem142

import (
	"unsafe"
	"github.com/HuihuangZhang/leetcode/common"
)

func detectCycle(head *common.ListNode[int]) *common.ListNode[int] {
	c := map[unsafe.Pointer]struct{}{}
	iter := head

	var cycleBeginNode *common.ListNode[int] = nil
	for iter != nil {
		if _, got := c[unsafe.Pointer(iter)]; got {
			cycleBeginNode = iter
			break
		}
		iter = iter.Next
	}
	return cycleBeginNode
}
