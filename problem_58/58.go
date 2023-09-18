package problem58

func reverseLeftWords(s string, n int) string {
	if len(s) == 0 || n % len(s) == 0 {
		return s
	}
	byteS := []byte(s)
	n = n % len(s)

	reverse(byteS[:n])
	reverse(byteS[n:])
	reverse(byteS)

	return string(byteS)
}

func reverse[t any](v []t) []t {
	left, right := 0, len(v) - 1
	for left < right {
		v[left], v[right] = v[right], v[left]
		left++
		right--
	}
	return v
}
