package solution

import (
	"reflect"
	"testing"
)

var tree = &TreeNode{
	Val: 5,
	Left: &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   11,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 2},
		},
	},
	Right: &TreeNode{
		Val:  8,
		Left: &TreeNode{Val: 13},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 1},
		},
	},
}

func TestPathSum(t *testing.T) {
	got := pathSum(tree, 22)
	want := [][]int{
		[]int{5, 4, 11, 2},
		[]int{5, 8, 4, 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, but got %v", want, got)
	}
}
