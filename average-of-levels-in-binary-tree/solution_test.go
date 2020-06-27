package solution

import (
	"fmt"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	r1 := &TreeNode{
		Val:   5,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: -3},
	}

	fmt.Println(averageOfLevels(r1))
}
