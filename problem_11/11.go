package problem_11

// Your runtime beats 27.3 % of golang submissions
// Your memory usage beats 17.01 % of golang submissions (8.4 MB)

func maxAreaV1(height []int) int {
	cursor_h, cursor_t := 0, len(height)-1
	max := 0
	for cursor_h < cursor_t {
		cur_area := (cursor_t - cursor_h) * min(height[cursor_t], height[cursor_h])
		if cur_area > max {
			max = cur_area
		}
		if height[cursor_h] > height[cursor_t] {
			cursor_t--
		} else {
			cursor_h++
		}
	}
	return max
}

func maxArea(height []int) int {
	cursor_h, cursor_t := 0, len(height)-1
	max := 0
	for cursor_h < cursor_t {
		cur_area := (cursor_t - cursor_h) * min(height[cursor_t], height[cursor_h])
		if cur_area > max {
			max = cur_area
		}
		if height[cursor_h] > height[cursor_t] {
			tmp := cursor_t
			cursor_t--
			for cursor_h < cursor_t && height[tmp] > height[cursor_t] {
				cursor_t--
			}
		} else {
			tmp := cursor_h
			cursor_h++

			for cursor_h < cursor_t && height[tmp] > height[cursor_h] {
				cursor_h++
			}
		}
	}
	return max
}

func min(l, r int) int {
	if l > r {
		return r
	}
	return l
}
