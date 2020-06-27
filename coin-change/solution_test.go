package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestCoinChanges(t *testing.T) {
	helper.AssertionSame(t, coinChange([]int{1, 2, 5}, 11), 3)
	helper.AssertionSame(t, coinChange([]int{2}, 3), -1)
	helper.AssertionSame(t, coinChange([]int{1}, 0), 0)
	helper.AssertionSame(t, coinChange([]int{186, 419, 83, 408}, 6249), 20)
}
