package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	head := treeToDoublyList(root.Left)
	if head == nil {
		head = root
	} else {
		tail := head
		for tail.Right != nil {
			tail = tail.Right
		}

		tail.Right = root
		root.Left = tail
	}

	rhead := treeToDoublyList(root.Right)
	if rhead != nil {
		root.Right = rhead
		rhead.Left = root
	}

	return head
}
