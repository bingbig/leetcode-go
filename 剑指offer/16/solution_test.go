package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestMyPow(t *testing.T) {
	helper.AssertionSame(t, myPow(2.0, 10), 1024)
	helper.AssertionSame(t, myPow(2.1, 3), 9.261)
	helper.AssertionSame(t, myPow(2, -2), 0.25)
}
