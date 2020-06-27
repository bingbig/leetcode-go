package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestNthUglyNumber(t *testing.T) {
	helper.AssertionSame(t, nthUglyNumber(10), 12)
}
