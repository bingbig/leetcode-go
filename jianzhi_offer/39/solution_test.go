package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestMajorityElement(t *testing.T) {
	helper.AssertionSame(t, majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}), 2)
	helper.AssertionSame(t, majorityElement([]int{1}), 1)
}
