package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestGetKthMagicNumber(t *testing.T) {
	helper.AssertionSame(t, getKthMagicNumber(5), 9)
}
