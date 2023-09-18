/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

package leetcode
import (
	// "fmt"
)

// @lc code=start

// 采用 DFS 的思路
func lexicalOrder(n int) []int {
	result := make([]int, 0, n)
	// 比较特殊的是，只有个位数是迭代 1-9，其余都是 0-9
	printLexical(1, 8, n, &result)	
	return result
}

func printLexical(anchor int, iterEnd int, max int, result *[]int) {
	for i := 0; i <= iterEnd && anchor + i <= max; i++ {
		*result = append(*result, anchor + i)
		if (anchor + i) * 10 <= max {
			printLexical((anchor + i) * 10, 9, max, result)
		}
	}
}
// @lc code=end

