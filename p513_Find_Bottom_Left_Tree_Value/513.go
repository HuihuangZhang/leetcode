package p513findbottomlefttreevalue

import (
	"github.com/HuihuangZhang/leetcode/common/tree"
)

func findBottomLeftValue(root *tree.Node[int]) int {
	if root == nil {
		return 0
	}

	queue := []*tree.Node[int]{}
	queue = append(queue, root)
	queue_tmp := []*tree.Node[int]{}
	result := 0
	
	for len(queue) != 0 {
		for i := 0; i < len(queue); i++ {
			if i == 0 {
				result = queue[0].Val
			}
			out := queue[i]
			if out.Left != nil {
				queue_tmp = append(queue_tmp, out.Left)
			}
			if out.Right != nil {
				queue_tmp = append(queue_tmp, out.Right)
			}
		}
		queue = nil
		queue = queue_tmp
		queue_tmp = []*tree.Node[int]{}
	}
	return result
}