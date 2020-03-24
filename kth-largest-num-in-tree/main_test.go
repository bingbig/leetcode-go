package main

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	if 0 != kthLargest(nil, 0) {
		t.Errorf("test empty root error")
	}

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}

	want := 1
	got := kthLargest(root, 4)
	if want != got {
		t.Errorf("test 2, want %d, but got %d", want, got)
	}

	root = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
	}

	want = 4
	got = kthLargest(root, 1)

	if want != got {
		t.Errorf("test 3, want %d, but got %d", want, got)
	}
}

func TestKthLargest2(t *testing.T) {
	want := 4
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 6},
	}
	got := kthLargest(root, 3)
	if want != got {
		t.Errorf("test 4, want %d, but got %d", want, got)
	}
}
