package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	rmax := len(matrix)
	if rmax == 0 {
		return []int{}
	}
	cmax := len(matrix[0])
	if cmax == 0 {
		return []int{}
	}

	output := make([]int, rmax*cmax)
	rmax--
	cmax--
	rmin := 0
	cmin := 0
	index := 0
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

func main() {
	m := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}

	fmt.Println(spiralOrder(m))

	m = [][]int{
		[]int{1, 2, 3, 4},
		[]int{5, 6, 7, 8},
		[]int{9, 10, 11, 12},
	}
	fmt.Println(spiralOrder(m))
}
