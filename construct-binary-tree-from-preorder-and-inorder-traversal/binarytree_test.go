package main

import "testing"

func TestPreorderString(t *testing.T) {
	tree := &TreeNode{Val: 3}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 20}
	tree.Right.Left = &TreeNode{Val: 15}
	tree.Right.Right = &TreeNode{Val: 7}

	want := "3\t9\t20\t15\t7\t"
	got := tree.PreorderString()
	if want != got {
		t.Errorf("want: %s, but got: %s", want, got)
	}
}
