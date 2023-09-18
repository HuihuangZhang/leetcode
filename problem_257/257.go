package problem257

import (
	"strconv"

	"github.com/HuihuangZhang/leetcode/common/tree"
)

const (
	pathSep = "->"
)

func binaryTreePaths(root *tree.Node[int]) []string {
	if root == nil {
		return []string{}
	}

	result := []string{}
	rootPath := strconv.FormatInt(int64(root.Val), 10)
	if root.Left != nil {
		result = append(result, getPath(root.Left, rootPath)...)
	}
	if root.Right != nil {
		result = append(result, getPath(root.Right, rootPath)...)
	}
	if root.Left == nil && root.Right == nil {
		result = append(result, rootPath)
	}
	return result
}

func getPath(node *tree.Node[int], parentPath string) []string {
	curPath := parentPath + pathSep + strconv.FormatInt(int64(node.Val), 10)

	resultPath := []string{}
	if node.Left != nil {
		resultPath = append(resultPath, getPath(node.Left, curPath)...)
	}
	if node.Right != nil {
		resultPath = append(resultPath, getPath(node.Right, curPath)...)
	}
	if node.Left == nil && node.Right == nil {
		resultPath = append(resultPath, curPath)
	}
	return resultPath
}
