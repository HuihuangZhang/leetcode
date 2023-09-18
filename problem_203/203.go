package problem203

import (
	"github.com/HuihuangZhang/leetcode/common"
)

func removeElements(head *common.ListNode[int], val int) *common.ListNode[int] {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return nil
	}

	prev := head
	iter := head.Next

	for iter != nil {
		if iter.Val == val {
			prev.Next = iter.Next
		} else {
			prev = prev.Next
		}

		iter = iter.Next
	}

	return head
}

