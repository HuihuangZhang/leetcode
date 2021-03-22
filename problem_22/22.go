package problem_22

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	result := helper(n)
	keys := []string{}
	for k, _ := range result {
		keys = append(keys, k)
	}
	return keys
}

func helper(n int) map[string]int {
	if n == 1 {
		return map[string]int{"()": 1}
	}

	tmp := helper(n - 1)
	result := map[string]int{}
	for k, _ := range tmp {
		result = checkAndAdd("()"+k, 1, result)
		result = checkAndAdd("("+k+")", 1, result)
		result = checkAndAdd(k+"()", 1, result)
	}
	return result
}

func checkAndAdd(key string, defaultV int, container map[string]int) map[string]int {
	if _, ok := container[key]; !ok {
		container[key] = defaultV
	}
	return container
}
