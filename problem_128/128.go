package problem128

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	container := map[int]bool{}
	var max, counter, start int

	max = 0
	for _, num := range nums {
		container[num] = true
	}
	for _, num := range nums {
		if container[num-1] {
			continue
		}
		counter = 1
		start = num + 1
		for {
			if container[start] {
				counter++
				start++
				continue // 如果下一个数还存在，即子序列还成立，继续找下一个
			} else if counter > max {
				max = counter
			}
			break
		}
	}
	return max
}
