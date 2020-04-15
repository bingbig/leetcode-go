package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		np := copyNode(p)
		p.Next = np
		p = np.Next
	}
	p = head
	nhead, np := head.Next, head.Next
	cn := head.Next
	for cn != nil {
		if cn.Random != nil {
			cn.Random = cn.Random.Next
		}
		if cn.Next != nil {
			cn = cn.Next.Next
		} else {
			break
		}
	}

	for np != nil {
		p.Next = np.Next
		p = p.Next
		if np.Next != nil {
			np.Next = np.Next.Next
		}
		np = np.Next
	}

	return nhead
}

func copyNode(n *Node) *Node {
	return &Node{
		Val:    n.Val,
		Next:   n.Next,
		Random: n.Random,
	}
}
