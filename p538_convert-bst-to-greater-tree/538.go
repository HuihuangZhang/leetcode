package p538convertbsttogreatertree

import "github.com/HuihuangZhang/leetcode/common/tree"

func convertBST(root *tree.Node[int]) *tree.Node[int] {
	if root == nil {
		return nil
	}
	greaterTree(root, 0)

	return root
}

func greaterTree(node *tree.Node[int], sum int) int {
	rightMax := sum
	if node.Right != nil {
		rightMax = greaterTree(node.Right, sum)
	}

	node.Val = rightMax + node.Val

	if node.Left != nil {
		return greaterTree(node.Left, node.Val)
	}

	return node.Val
}
