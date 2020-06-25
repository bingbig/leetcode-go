package solution

import (
	"fmt"
	"testing"
)

func TestSubIncrSequence(t *testing.T) {
	nums := []int{3, 4, 5, 7, 2, 3, 5, 47, 90, 90, 32, 2, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(subIncrSequence(nums))
}
