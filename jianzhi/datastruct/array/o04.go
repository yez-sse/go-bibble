package array

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	// for还可以当其他语言里的while来用，神奇
	for i < len(matrix) && j > -1 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
