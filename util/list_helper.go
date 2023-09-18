package util

import "github.com/HuihuangZhang/leetcode/common"

func Arr2List(arr []int) *common.ListNode[int] {
	if len(arr) == 0 {
		return nil
	}
	head := &common.ListNode[int]{
		Val: arr[0],
	}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &common.ListNode[int]{
			Val: arr[i],
		}
		cur = cur.Next
	}
	return head
}

func List2Arr(list *common.ListNode[int]) []int {
	arr := []int{}
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}
