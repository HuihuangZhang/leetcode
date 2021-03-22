package problem_54

type Direction int

// enum 类型
const (
	East  Direction = iota
	West  Direction = iota
	North Direction = iota
	South Direction = iota
)

func spiralOrder(matrix [][]int) []int {
	directions := []Direction{East, South, West, North}

	width, length := len(matrix), len(matrix[0])

	result := []int{}
	var direcHelper int8 = 0
	// 通过不断调整开始、终止行列来压缩
	iStart, iEnd, jStart, jEnd := 0, width-1, 0, length-1
	i := iStart
	j := jStart

	count := width * length
	for {
		switch directions[direcHelper%4] {
		case East:
			for j = jStart; j <= jEnd; j++ {
				result = append(result, matrix[i][j])
			}
			j--
			iStart++
			direcHelper++
		case South:
			for i = iStart; i <= iEnd; i++ {
				result = append(result, matrix[i][j])
			}
			i--
			jEnd--
			direcHelper++
		case West:
			for j = jEnd; j >= jStart; j-- {
				result = append(result, matrix[i][j])
			}
			j++
			iEnd--
			direcHelper++
		case North:
			for i = iEnd; i >= iStart; i-- {
				result = append(result, matrix[i][j])
			}
			i++
			jStart++
			direcHelper++
		}
		if len(result) == count {
			return result
		}
	}
}
