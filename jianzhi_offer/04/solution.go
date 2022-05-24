package solution

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	rows := len(matrix)
	cols := len(matrix[0])
	r, c := 0, cols-1
	for {
		if r >= rows || c < 0 {
			return false
		}
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			r++
		} else {
			c--
		}
	}
}
