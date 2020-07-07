package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestFindMagicIndex(t *testing.T) {
	arr := []int{0, 2, 3, 4, 5}
	helper.AssertionSame(t, findMagicIndex(arr), 0)
	helper.AssertionSame(t, findMagicIndex([]int{1, 1, 1}), 1)
	helper.AssertionSame(t, findMagicIndex([]int{-1, 0, 1, 3, 4, 5, 6}), 3)
}
