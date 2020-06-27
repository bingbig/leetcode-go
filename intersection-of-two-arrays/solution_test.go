package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestIntersection(t *testing.T) {
	helper.AssertionSame(t, intersection([]int{1, 2, 2, 1}, []int{2, 2}), []int{2})
	helper.AssertionSame(t, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), []int{9, 4})
}
