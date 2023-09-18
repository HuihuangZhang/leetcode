package problem160

import (
	"github.com/HuihuangZhang/leetcode/common"
)

func getIntersectionNode(headA, headB *common.ListNode[int]) *common.ListNode[int] {
	lenA, lenB := calcListLength(headA), calcListLength(headB)

	var long, short *common.ListNode[int]
	cntInAdvance := 0

	if lenA > lenB {
		long = headA
		short = headB
		cntInAdvance = lenA - lenB
	} else {
		short = headA
		long = headB
		cntInAdvance = lenB - lenA
	}

	for cntInAdvance != 0 {
		long = long.Next
		cntInAdvance--
	}

	var intersectionNode *common.ListNode[int] = nil
	for long != nil && short != nil {
		if long == short {
			intersectionNode = long
			break
		}
		long = long.Next
		short = short.Next
	}
	return intersectionNode
}

func calcListLength(head *common.ListNode[int]) int {
	iter := head
	cnt := 0
	for iter != nil {
		cnt++
		iter = iter.Next
	}
	return cnt
}