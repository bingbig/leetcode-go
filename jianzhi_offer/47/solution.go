package solution

func maxValue(grid [][]int) int {
	if grid == nil || len(grid) == 0 || grid[0] == nil || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			max := 0
			if r-1 >= 0 {
				max = MaxInt(max, grid[r-1][c])
			}
			if c-1 >= 0 {
				max = MaxInt(max, grid[r][c-1])
			}
			grid[r][c] += max
		}
	}

	return grid[rows-1][cols-1]
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
