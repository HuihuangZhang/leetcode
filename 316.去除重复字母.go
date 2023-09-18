package leetcode

/*
 * @lc app=leetcode.cn id=316 lang=golang
 *
 * [316] 去除重复字母
 */

// @lc code=start

import (
	"strings"
)

func removeDuplicateLetters(s string) string {
	flag := [26]int{}
	// 统计所有字母出现的次数
	for _, char := range s {
		flag[int(char - 'a')]++
	}

	// 开始用单调栈的思路
	monoStack := make([]int, 26)
	stackTop := -1 
	isInStack := [26]bool{}

	for _, char := range s {
		intChar := int(char - 'a')	
		if isInStack[intChar] {
			flag[intChar]--
			continue
		}
		for stackTop >= 0 {
			if monoStack[stackTop] >= intChar && flag[monoStack[stackTop]] > 0 {
				isInStack[monoStack[stackTop]] = false
				stackTop--
				continue
			}
			break
		}	
		stackTop++
		monoStack[stackTop] = intChar
		flag[intChar]--
		isInStack[intChar] = true
	}

	buf := strings.Builder{}
	for i := 0; i <= stackTop; i++ {
		buf.WriteRune(rune(monoStack[i] + 'a'))
	}

	return buf.String()
}
// @lc code=end

