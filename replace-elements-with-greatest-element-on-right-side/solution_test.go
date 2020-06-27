package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestReplaceElements(t *testing.T) {
	helper.AssertionSame(t, replaceElements([]int{17, 18, 5, 4, 6, 1}), []int{18, 6, 6, 6, 1, -1})
}
