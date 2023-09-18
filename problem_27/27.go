package problem27

func removeInPlace(arr []int, num int) (int, []int) {
	if len(arr) == 0 {
		return 0, arr
	}
	i, j := 0, 1
	for j < len(arr) {
		if arr[i] != num {
			i++
			j++
		} else {
			for j < len(arr) && arr[j] == num {
				j++
			}
			// 如果找不到可以替换的，直接退出
			if j < len(arr) {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j++
			}
		}
	}

	if arr[i] == num {
		return i, arr
	}

	return i + 1, arr
}

func removeInPlaceV2(arr []int, num int) (int, []int) {
	lowPtr, fastPtr := 0, 0
	for fastPtr = 0; fastPtr < len(arr); fastPtr++ {
		if arr[fastPtr] != num {
			arr[lowPtr] = arr[fastPtr]
			lowPtr++
		}
	}
	return lowPtr, arr
}
