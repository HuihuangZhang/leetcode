package problem145

import (
	"github.com/HuihuangZhang/leetcode/common/tree"
)

func postorderTraversal(root *tree.Node[int]) []int {
	if root == nil {
		return []int{}
	}

	container := []int{}
	container = append(container, postorderTraversal(root.Left)...)
	container = append(container, postorderTraversal(root.Right)...)
	container = append(container, root.Val)
	return container
}