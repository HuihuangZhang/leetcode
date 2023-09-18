package problem206

import (
	"github.com/HuihuangZhang/leetcode/common"
)

func reverseList(head *common.ListNode[int]) *common.ListNode[int] {
	if head == nil {
		return head
	}

	var cur, next  *common.ListNode[int]
	cur, next = nil, head

	for next != nil {
		tmp := next.Next

		next.Next = cur
		cur = next
		next = tmp
	}
	return cur
}