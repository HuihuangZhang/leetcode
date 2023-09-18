package problem76

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	
	letterCnt := [256]int{}
	totalCnt := 0
	cnt := 0
	result := ""

	for _, bval := range []byte(t) {
		letterCnt[bval]++
		totalCnt++
	}

	left, right := 0, 0
	for ;right < len(s); right++ {
		letterCnt[s[right]]--
		if letterCnt[s[right]] >= 0 {
			cnt++
		}

		for cnt == totalCnt {
			letterCnt[s[left]]++
			if letterCnt[s[left]] > 0 {
				if len(result) == 0 || len(result) > (right - left + 1) {
					result = s[left:right + 1]
				}
				cnt--
			}
			left++
		}
	}
	return result
}
