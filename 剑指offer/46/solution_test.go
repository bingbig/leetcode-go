package solution

import (
	"testing"

	"github.com/bingbig/leetcode/helper"
)

func TestTranslateNum(t *testing.T) {
	helper.AssertionSame(t, translateNum(12258), 5)
}
