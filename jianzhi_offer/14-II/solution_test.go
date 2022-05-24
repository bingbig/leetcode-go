package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestCuttingRope(t *testing.T) {
	helper.AssertionSame(t, cuttingRope(10), 36)
	helper.AssertionSame(t, cuttingRope(120), 953271190)
}
