package problem404

import "github.com/HuihuangZhang/leetcode/common/tree"

func sumOfLeftLeaves(root *tree.Node[int]) int {
	if root == nil {
		return 0
	}
	return sumOfLeavesRecursion(root, false)
}

func sumOfLeavesRecursion(node *tree.Node[int], isLeft bool) int {
	if node.Left == nil && node.Right == nil {
		if isLeft {
			return node.Val
		} else {
			return 0
		}
	}

	sum := 0
	if node.Left != nil {
		sum += sumOfLeavesRecursion(node.Left, true)
	}
	if node.Right != nil {
		sum += sumOfLeavesRecursion(node.Right, false)
	}
	return sum
}