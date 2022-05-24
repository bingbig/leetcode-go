package solution

import (
	"fmt"
	"testing"
)

var tree = &TreeNode{
	Val:  3,
	Left: &TreeNode{Val: 9},
	Right: &TreeNode{
		Val:   20,
		Left:  &TreeNode{Val: 15},
		Right: &TreeNode{Val: 7},
	},
}

func TestLevelOrder(t *testing.T) {
	fmt.Println(levelOrder(tree))

}
