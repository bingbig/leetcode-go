package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	helper.AssertionSame(t, lengthOfLongestSubstring("abcabcbb"), 3)
	helper.AssertionSame(t, lengthOfLongestSubstring("bbbbbb"), 1)
	helper.AssertionSame(t, lengthOfLongestSubstring("pwwkew"), 3)
	helper.AssertionSame(t, lengthOfLongestSubstring("aab"), 2)
	helper.AssertionSame(t, lengthOfLongestSubstring("dvdf"), 3)
}
