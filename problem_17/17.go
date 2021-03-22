package problem_17

var vocab = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := []string{}
	if len(digits) == 0 {
		return result
	}

	result = vocab[digits[0]-'0'-2]
	for i := 1; i < len(digits); i++ {
		tmp := []string{}
		index := digits[i] - '0' - 2

		for _, j := range vocab[index] {
			for _, k := range result {
				tmp = append(tmp, k+j)
			}
		}

		result = tmp
	}
	return result
}
