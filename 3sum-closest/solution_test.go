package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestThreeSumClosest(t *testing.T) {
	helper.AssertionSame(t, threeSumClosest([]int{-1, 2, 1, -4}, 1), 2)
	helper.AssertionSame(t, threeSumClosest([]int{1, 1, 1, 1}, 1), 3)
}
