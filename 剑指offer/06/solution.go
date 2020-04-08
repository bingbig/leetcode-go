package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	rtn := []int{}
	p := head
	for p != nil {
		rtn = append(rtn, p.Val)
		p = p.Next
	}

	for i, j := 0, len(rtn)-1; i < j; {
		rtn[i], rtn[j] = rtn[j], rtn[i]
		i++
		j--
	}

	return rtn
}

/**
 * 执行用时 :56 ms, 在所有 Go 提交中击败了7.89%的用户
 * 内存消耗 :9.5 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func reversePrint2(head *ListNode) []int {
	if head == nil {
		return nil
	}

	rtn := []int{}
	p := head
	for p != nil {
		rtn = append([]int{p.Val}, rtn...)
		p = p.Next
	}

	return rtn
}
