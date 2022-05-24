package solution

func spiralOrder(matrix [][]int) []int {
	rmax := len(matrix)
	if rmax == 0 {
		return []int{}
	}

	cmax := len(matrix[0])
	if cmax == 0 {
		return []int{}
	}

	output := make([]int, cmax*rmax)

	rmin, cmin, index := 0, 0, 0
	cmax--
	rmax--
	for rmin <= rmax && cmin <= cmax {
		// upper row
		for i := cmin; i <= cmax; i++ {
			output[index] = matrix[rmin][i]
			index++
		}

		// right col
		for j := rmin + 1; j <= rmax; j++ {
			output[index] = matrix[j][cmax]
			index++
		}

		// bottom row
		for i := cmax - 1; i >= cmin && rmin < rmax; i-- {
			output[index] = matrix[rmax][i]
			index++
		}

		// left col
		for j := rmax - 1; j > rmin && cmax > cmin; j-- {
			output[index] = matrix[j][cmin]
			index++
		}
		rmin++
		cmin++
		rmax--
		cmax--
	}

	return output
}
