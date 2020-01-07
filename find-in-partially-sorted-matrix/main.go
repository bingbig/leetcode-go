package main

import (
	"fmt"
)

func Find(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows < 1 {
		return false
	}
	cols := len(matrix[0])
	if cols < 1 {
		return false
	}

	i := 0
	j := cols - 1
	for i < rows && j >= 0 {
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

func main() {
	matrix := [][]int{
		[]int{1, 2, 8, 9},
		[]int{2, 4, 9, 12},
		[]int{4, 7, 10, 13},
		[]int{6, 8, 11, 15},
	}

	fmt.Println(Find(matrix, 7))
	fmt.Println(Find(matrix, 9))
	fmt.Println(Find(matrix, 100))
	fmt.Println(Find(matrix, 10))
}
