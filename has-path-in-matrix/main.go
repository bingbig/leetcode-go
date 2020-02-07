// 剑指offer 面试题12 矩阵中的路径，p89
package main

import "fmt"

func hasPath(matrix [][]byte, path []byte) bool {
	if len(path) == 0 {
		return true
	}

	rows := len(matrix)
	if rows == 0 {
		return false
	}

	cols := len(matrix[0])
	visited := make([]bool, cols*rows)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if hasPathCore(matrix, i, j, rows, cols, path, visited) {
				return true
			}
		}
	}

	return false
}

func hasPathCore(matrix [][]byte, row, col, rows, cols int, path []byte, visited []bool) bool {
	flag := false
	if row < rows && col < cols && row >= 0 && col >= 0 && !visited[row*cols+col] && matrix[row][col] == path[0] {
		if len(path) <= 1 {
			return true
		}

		visited[row*cols+col] = true
		flag = hasPathCore(matrix, row, col-1, rows, cols, path[1:], visited) ||
			hasPathCore(matrix, row, col+1, rows, cols, path[1:], visited) ||
			hasPathCore(matrix, row-1, col, rows, cols, path[1:], visited) ||
			hasPathCore(matrix, row+1, col, rows, cols, path[1:], visited)

		if !flag {
			visited[row*cols+col] = false
		}
	}

	return flag
}

func main() {
	m1 := [][]byte{
		[]byte{'a', 'b', 't', 'g'},
		[]byte{'c', 'f', 'c', 's'},
		[]byte{'j', 'd', 'e', 'h'},
	}

	path := []byte{'b', 'f', 'c', 'e'}
	fmt.Println(hasPath(m1, path))                       // true
	fmt.Println(hasPath(m1, []byte{'a', 'b', 'f', 'b'})) // false
}
