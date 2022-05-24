package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestCountDigitOne(t *testing.T) {
	helper.AssertionSame(t, countDigitOne(10), 2)
	helper.AssertionSame(t, countDigitOne(12), 5)
}
