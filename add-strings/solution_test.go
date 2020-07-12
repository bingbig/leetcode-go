package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestAddStrings(t *testing.T) {
	helper.AssertionSame(t, addStrings("12345", "98765"), "111110")
	helper.AssertionSame(t, addStrings("9", "99"), "108")
	helper.AssertionSame(t, addStrings("", "9"), "9")
	helper.AssertionSame(t, addStrings("98", "9"), "107")
}
