package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestSmallestRangeII(t *testing.T) {
	helper.AssertionSame(t, smallestRangeII([]int{1, 3, 6}, 3), 3)
}
