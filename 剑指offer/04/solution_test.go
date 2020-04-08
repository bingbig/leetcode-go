package solution

import "testing"

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		[]int{1, 4, 7, 11, 15},
		[]int{2, 5, 8, 12, 19},
		[]int{3, 6, 9, 16, 22},
		[]int{10, 13, 14, 17, 24},
		[]int{18, 21, 23, 26, 30},
	}
	got := findNumberIn2DArray(matrix, 5)
	if false == got {
		t.Errorf("assert true but got false")
	}
	got = findNumberIn2DArray(matrix, 20)
	if true == got {
		t.Errorf("assert false but got true")
	}
}
