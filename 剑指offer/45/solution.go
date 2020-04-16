package solution

import (
	"bytes"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	numStr := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numStr[i] = strconv.Itoa(nums[i])
	}

	sort.Slice(numStr, func(i, j int) bool {
		return numStr[j]+numStr[i] > numStr[i]+numStr[j]
	})
	var buf bytes.Buffer
	for i := range numStr {
		buf.WriteString(numStr[i])
	}

	return buf.String()
}
