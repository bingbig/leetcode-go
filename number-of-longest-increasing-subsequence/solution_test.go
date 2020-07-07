package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestFindNumberOfLIS(t *testing.T) {
	helper.AssertionSame(t, findNumberOfLIS([]int{1, 3, 5, 4, 7}), 2)
	helper.AssertionSame(t, findNumberOfLIS([]int{2, 2, 2, 2, 2}), 5)
}
