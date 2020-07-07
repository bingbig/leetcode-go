package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestFindKthLargest(t *testing.T) {
	helper.AssertionSame(t, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	helper.AssertionSame(t, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
}
