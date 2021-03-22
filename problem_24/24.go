package problem_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	cursor := head
	newHead := head.Next
	for cursor != nil {
		if cursor.Next == nil {
			break
		} else {
			swapTwoNode(prev, cursor, cursor.Next)
		}
		// 因为位置换了，所以这里只要指向下一个就可以了
		prev = cursor
		cursor = cursor.Next
	}
	return newHead
}

func swapTwoNode(prev *ListNode, first *ListNode, second *ListNode) {
	if prev != nil {
		prev.Next = second
	}
	first.Next = second.Next
	second.Next = first
}

func arr2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{
		Val: arr[0],
	}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val: arr[i],
		}
		cur = cur.Next
	}
	return head
}

func list2Arr(list *ListNode) []int {
	arr := []int{}
	for list != nil {
		arr = append(arr, list.Val)
		list = list.Next
	}
	return arr
}
