package problem144

import (
	"github.com/HuihuangZhang/leetcode/common/stack"
	"github.com/HuihuangZhang/leetcode/common/tree"
)

func preorderTraversal(root *tree.Node[int]) []int {
	if root == nil {
		return []int{}
	}
	container := []int{}

	container = append(container, root.Val)
	container = append(container, preorderTraversal(root.Left)...)
	container = append(container, preorderTraversal(root.Right)...)
	return container
}

func preorderTraversalIteration(root *tree.Node[int]) []int {
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
		result = append(result, cursor.Val)
		if cursor != nil {
			_stack.Push(cursor.Right)
			_stack.Push(cursor.Left)
		}
	}
	return result
}
