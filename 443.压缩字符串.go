/*
 * @lc app=leetcode.cn id=443 lang=golang
 *
 * [443] 压缩字符串
 */

package leetcode

// import "fmt"

// @lc code=start
func compress(chars []byte) int {
	flag := [256]uint16{}

	previewChar := chars[0]
	flag[previewChar]++

	resultCnt := 0
	
	for i := 1; i < len(chars); i++ {
		if chars[i] != previewChar {
			resultCnt = overwriteChars(&chars, resultCnt, flag, int(previewChar))

			flag[previewChar] = 0
			previewChar = chars[i]
		}
		flag[chars[i]]++
	}
	// 最后一个字符
	resultCnt = overwriteChars(&chars, resultCnt, flag, int(previewChar))

	return resultCnt
}

func overwriteChars(chars *[]byte, overwriteIdx int, flag [256]uint16, char int) int {
	if flag[char] == 1 {
		(*chars)[overwriteIdx] = byte(char)
		overwriteIdx++ // 原字符
	} else {
		(*chars)[overwriteIdx] = byte(char)
		overwriteIdx++ // 原字符
		cnt := flag[char]
		// 统计字符出现的次数

		num := make([]byte, 0, 4)
		for cnt != 0 {
			num = append(num, uint8(cnt % 10) + '0')
			cnt = cnt / 10
		}
		for i := len(num) - 1; i >= 0; i-- {
			(*chars)[overwriteIdx] = num[i] 
			overwriteIdx++ 
		}

	}
	return overwriteIdx 
}
// @lc code=end

