package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestIsMatch(t *testing.T) {
	helper.AssertionSame(t, isMatch("aa", "a"), false)
	helper.AssertionSame(t, isMatch("aa", "a*"), true)
	helper.AssertionSame(t, isMatch("ab", ".*"), true)
	helper.AssertionSame(t, isMatch("aab", "c*a*b"), true)
	helper.AssertionSame(t, isMatch("mississippi", "mis*is*p*."), false)
	helper.AssertionSame(t, isMatch("a", ".*..a*"), false)
}
