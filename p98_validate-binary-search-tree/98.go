package p98validatebinarysearchtree

import (
	"math"

	"github.com/HuihuangZhang/leetcode/common/tree"
)

func isValidBST(root *tree.Node[int]) bool {
	if root == nil {
		return true
	}
	return isTreeInRange(root, math.MinInt32, math.MaxInt32)
}

func isTreeInRange(node *tree.Node[int], minVal, maxVal int) bool {
	isNodeValid := minVal < node.Val && node.Val < maxVal

	if !isNodeValid {
		return false
	}

	if node.Left != nil {
		isLeftTreeValid := isTreeInRange(node.Left, minVal, node.Val)
		if !isLeftTreeValid {
			return false
		}
	}

	if node.Right != nil {
		isRightTreeValid := isTreeInRange(node.Right, node.Val, maxVal)
		if !isRightTreeValid {
			return false
		}
	}

	return true
}