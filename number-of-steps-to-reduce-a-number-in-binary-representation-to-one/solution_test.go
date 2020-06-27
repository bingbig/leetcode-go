package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestNumSteps(t *testing.T) {
	helper.AssertionSame(t, numSteps("1101"), 6)
	helper.AssertionSame(t, numSteps("10"), 1)
}
