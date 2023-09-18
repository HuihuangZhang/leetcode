package problem209


func minSubArrayLen(arr []int, s int) int {
	i, j := 0, 0

	sum := 0
	minLen := 0

	for j <= len(arr) && i <= j {
		if sum >= s {
			if minLen == 0 || (j - i) < minLen {
				minLen = j - i
			}
			sum -= arr[i]
			i++
		} else {
			if j < len(arr) {
				sum += arr[j]
			}
			j++
		}
	}
	return minLen
}