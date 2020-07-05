package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestMinSubArrayLen(t *testing.T) {
	helper.AssertionSame(t, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}), 2)
	helper.AssertionSame(t, minSubArrayLen(7, []int{}), 0)
	helper.AssertionSame(t, minSubArrayLen(4, []int{1, 4, 4}), 1)
	helper.AssertionSame(t, minSubArrayLen(11, []int{1, 2, 3, 4, 5}), 3)
	helper.AssertionSame(t, minSubArrayLen(3, []int{1, 1}), 0)
	helper.AssertionSame(t, minSubArrayLen(15, []int{1, 2, 3, 4, 5}), 5)
}
