package solution

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	AssertionSame(t, lengthOfLongestSubstring("abcabcbb"), 3)
	AssertionSame(t, lengthOfLongestSubstring("bbbbbb"), 1)
	AssertionSame(t, lengthOfLongestSubstring("pwwkew"), 3)
	AssertionSame(t, lengthOfLongestSubstring("aab"), 2)
	AssertionSame(t, lengthOfLongestSubstring("dvdf"), 3)
}

func AssertionSame(t *testing.T, actual, expect interface{}) {
	switch actual.(type) {
	case int:
		if actual.(int) != expect.(int) {
			t.Errorf("want %d, but got %d", expect.(int), actual.(int))
		}
	}
}
