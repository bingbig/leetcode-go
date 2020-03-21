package main

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	stackA, stackB := list.New(), list.New()
	for pa != nil {
		stackA.PushBack(pa)
		pa = pa.Next
	}
	for pb != nil {
		stackB.PushBack(pb)
		pb = pb.Next
	}

	ta, tb := stackA.Back().Value.(*ListNode), stackB.Back().Value.(*ListNode)

	for ta == tb {
		stackA.Remove(stackA.Back())
		stackB.Remove(stackB.Back())
		tmp := ta
		ta = stackA.Back().Value.(*ListNode)
		tb = stackB.Back().Value.(*ListNode)

		if stackA.Len() == 0 || stackB.Len() == 0 {
			if ta == tb {
				return ta
			} else {
				return tmp
			}
		}

		if ta != tb {
			return tmp
		}
	}

	return nil
}
