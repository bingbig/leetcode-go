package solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	token []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.token = strings.Split(data, ",")
	return this.dfsDeserialize()
}

func (this *Codec) dfsDeserialize() *TreeNode {
	node := this.token[0]
	this.token = this.token[1:]
	if node == "#" {
		return nil
	}

	v, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   v,
		Left:  this.dfsDeserialize(),
		Right: this.dfsDeserialize(),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
