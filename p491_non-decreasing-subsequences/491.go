package p491nondecreasingsubsequences

var (
	result [][]int
	path []int
)

func findSubsequences(nums []int) [][]int {
	result = [][]int{}
	path = []int{}

	backtracing(nums, 0)
	return result
}

func backtracing(nums []int, startIndex int) {
	if len(path) >= 2 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		result = append(result, tmp)
	}

	used := [201]int{}
	for i := startIndex; i < len(nums); i++ {
		if (len(path) > 0 && nums[i] < path[len(path) - 1]) || used[nums[i] + 100] == 1 {
			continue
		}
		used[nums[i] + 100] = 1
		path = append(path, nums[i])
		backtracing(nums, i + 1)
		path = path[:len(path) - 1]
	}
}
