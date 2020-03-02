package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	tree := strconv.Itoa(root.Val)
	stack := list.New()
	stack.PushBack(root.Right)
	node := root.Left

	for stack.Len() > 0 {
		if node == nil {
			tree += ",n"
			node = stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
		} else {
			tree += "," + strconv.Itoa(node.Val)
			stack.PushBack(node.Right)
			node = node.Left
		}
	}

	if node == nil {
		tree += ",n"
	} else {
		tree += "," + strconv.Itoa(node.Val) + ",n,n"
	}

	return tree
}

func deserilize(tree string) *TreeNode {
	if tree == "" || tree[0] == 'n' {
		return nil
	}

	strs := strings.Split(tree, ",")
	root := strToTreeNode(strs[0])
	stack := list.New()
	stack.PushBack(root)
	node := root
	idx := 1
	for stack.Len() > 0 && idx < len(strs) {
		nnode := strToTreeNode(strs[idx])
		if node == nil {
			pnode := stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
			pnode.Right = nnode
			node = nnode
		} else {
			stack.PushBack(node)
			node.Left = nnode
			node = node.Left
		}
		idx++
	}

	return root
}

func strToTreeNode(str string) *TreeNode {
	if str == "n" {
		return nil
	}

	if val, err := strconv.Atoi(str); err == nil {
		return &TreeNode{Val: val}
	} else {
		return nil
	}
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 1},
			// Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}
	fmt.Println(serialize(root))
	fmt.Println(serialize(deserilize(serialize(root))))
}
