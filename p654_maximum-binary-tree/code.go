package p654maximumbinarytree

import "github.com/HuihuangZhang/leetcode/common/tree"


func constructMaximumBinaryTree(nums []int) *tree.Node[int] {
	if len(nums) == 0 {
		return nil
	}
	rootIdx := findMaxIdx(nums)
	rootNode := tree.NewNode(nums[rootIdx], nil, nil)
	constructHelper(nums[:rootIdx], rootNode, true)
	constructHelper(nums[rootIdx+1:], rootNode, false)
	return rootNode
}

func constructHelper(nums []int, parent *tree.Node[int], isLeft bool) {
	if len(nums) == 0 {
		return
	} else if len(nums) == 1 {
		if isLeft {
			parent.Left = tree.NewNode(nums[0], nil, nil)
		} else { 
			parent.Right = tree.NewNode(nums[0], nil, nil)
		}
		return
	}

	pivot := findMaxIdx(nums)
	childNode := tree.NewNode(nums[pivot], nil, nil)
	if isLeft {
		parent.Left = childNode
	} else { 
		parent.Right = childNode
	}

	constructHelper(nums[:pivot], childNode, true)
	constructHelper(nums[pivot+1:], childNode, false)
}


func findMaxIdx(nums []int) int {
	idx, max := 0, -1

	for i, num := range nums {
		if num > max {
			idx = i
			max = num
		}
	}
	return idx
}
