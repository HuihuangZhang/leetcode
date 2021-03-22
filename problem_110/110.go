package problem_110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := countHeight(root.Left)
	r := countHeight(root.Right)

	if math.Abs(float64(l-r)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalancedV2(root *TreeNode) bool {
	b, _ := checkAndCount(root)
	return b
}

func checkAndCount(tree *TreeNode) (bool, int) {
	if tree.Left == nil && tree.Right == nil {
		return true, 1 // 该节点为叶子节点
	}
	lb, rb, lh, rh := true, true, 0, 0
	if tree.Left != nil {
		lb, lh = checkAndCount(tree.Left)
	}
	// 在左子树平衡的条件下再判断右子树
	if tree.Right != nil && lb {
		rb, rh = checkAndCount(tree.Right)
	}

	isBalance := lb && rb && math.Abs(float64(lh-rh)) < 1
	height := max(lh, rh) + 1
	return isBalance, height
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func countHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}

	return max(countHeight(t.Left), countHeight(t.Right)) + 1
}

func constructTree(arr *[]*int) *TreeNode {
	nodeArr := make([]*TreeNode, len(*arr))
	for i, v := range *arr {
		if v != nil {
			nodeArr[i] = &TreeNode{
				Val:   *v,
				Left:  nil,
				Right: nil,
			}
		}
	}
	return makeTree(nodeArr, len(*arr), 0)
}

func makeTree(rt []*TreeNode, length int, idx int) *TreeNode {
	if idx >= length {
		return rt[0]
	}
	item := rt[idx]
	if item == nil {
		return nil
	}
	l := 2*idx + 1
	r := 2*idx + 2

	if l < length {
		item.Left = rt[l]
	}
	if r < length {
		item.Right = rt[r]
	}
	makeTree(rt, length, l)
	makeTree(rt, length, r)
	return rt[0]
}
