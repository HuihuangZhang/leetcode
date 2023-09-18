package problem977

func SquareOrderedList(input []int) []int {
	if len(input) == 0 {
		return input
	}
	pivot := findPivot(input, 0)

	result := make([]int, 0, len(input))
	negtivePtr, nonnegtivePtr := pivot - 1, pivot
	for negtivePtr >= 0 && nonnegtivePtr <= len(input) - 1 {
		if (-input[negtivePtr]) < input[nonnegtivePtr] {
			result = append(result, input[negtivePtr] * input[negtivePtr])
			negtivePtr--
		} else {
			result = append(result, input[nonnegtivePtr] * input[nonnegtivePtr])
			nonnegtivePtr++
		}
	}
	if negtivePtr < 0 {
		for nonnegtivePtr < len(input) {
			result = append(result, input[nonnegtivePtr] * input[nonnegtivePtr])
			nonnegtivePtr++
		}
	} else if nonnegtivePtr > len(input) - 1 {
		for negtivePtr >= 0 {
			result = append(result, input[negtivePtr] * input[negtivePtr])
			negtivePtr--
		}
	}
	return result
}

func findPivot(input []int, value int) int {
	left, right := 0, len(input) - 1

	for left <= right {
		middle := left + ((right - left) >> 1)

		if input[middle] == value {
			return middle
		} else if input[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return left
}
