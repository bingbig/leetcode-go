package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	helper.AssertionSame(t, spiralOrder(matrix), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})

	matrix = [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}
	helper.AssertionSame(t, spiralOrder(matrix), []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7})
}
