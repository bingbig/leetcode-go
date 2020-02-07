// https://leetcode-cn.com/problems/search-a-2d-matrix/
package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	if cols == 0 {
		return false
	}

	min := 0
	max := rows*cols - 1
	mid := rows * cols / 2

	for min <= max {
		midv := matrix[mid/cols][mid%cols]
		if midv == target || target == matrix[min/cols][min%cols] || target == matrix[max/cols][max%cols] {
			return true
		} else if midv > target {
			max = mid - 1
			min++
		} else if midv < target {
			min = mid + 1
			max--
		}
	}

	return false
}

func main() {
	m1 := [][]int{
		[]int{1, 3, 5, 7},
		[]int{10, 11, 16, 20},
		[]int{23, 30, 34, 50},
	}

	fmt.Println(searchMatrix(m1, 3))
	fmt.Println(searchMatrix(m1, 13))
	fmt.Println(searchMatrix(m1, 20))

	m2 := [][]int{[]int{1}}
	fmt.Println(searchMatrix(m2, 1))
}
