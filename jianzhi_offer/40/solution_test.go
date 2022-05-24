package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestGetLeastNumber(t *testing.T) {
	helper.AssertionSame(t, getLeastNumbers([]int{3, 2, 1}, 2), []int{1, 2})
	helper.AssertionSame(t, getLeastNumbers([]int{0, 1, 2, 1}, 1), []int{0})
}

func TestQuickSort(t *testing.T) {
	helper.AssertionSame(t, quickSort([]int{3, 2, 1}), []int{1, 2, 3})
}
