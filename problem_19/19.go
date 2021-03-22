package problem_19

import (
	"github.com/HuihuangZhang/leetcode/common"
)

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	iter := head

	for i := 0; i < n; i++ {
		iter = iter.Next
	}
	// 如果直接等于 nil，则说明删除的是第一个元素
	if iter == nil {
		return head.Next
	}
	prev := head
	for iter.Next != nil {
		prev = prev.Next
		iter = iter.Next
	}

	prev.Next = prev.Next.Next
	return head
}
