package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length; i++ {
		for j := 0; j < length/2; j++ {
			matrix[i][j], matrix[i][length-1-j] = matrix[i][length-1-j], matrix[i][j]
		}
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			matrix[i][j], matrix[length-j-1][length-i-1] = matrix[length-j-1][length-i-1], matrix[i][j]
		}
	}
}

func printImg(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(col)
			fmt.Print("\t")
		}
		fmt.Print("\n")
	}
}

func main() {
	m := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	printImg(m)
	rotate(m)
	printImg(m)
}
