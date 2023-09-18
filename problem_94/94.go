package problem94

import (
	"github.com/HuihuangZhang/leetcode/common/tree"
	"github.com/HuihuangZhang/leetcode/common/stack"
)

func inorderTraversalIteration(root *tree.Node[int]) []int {
	if root == nil {
		return []int{}
	}
	_stack := stack.Init[*tree.Node[int]](100)

	_stack.Push(root)
	result := []int{}

	for !_stack.IsEmpty() {
		cursor := _stack.Pop()
		if cursor == nil {
			continue
		}

		if cursor.Left == nil && cursor.Right == nil {
			result = append(result, cursor.Val)
			continue
		}
		l, r := cursor.Left, cursor.Right
		cursor.Left, cursor.Right = nil, nil

		_stack.Push(r)
		_stack.Push(cursor)
		_stack.Push(l)
	}

	return result
}
