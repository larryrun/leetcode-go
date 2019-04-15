package q2

/*
https://leetcode.com/problems/add-two-numbers/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var firstNode, lastNode *ListNode
	carry := 0
	for {
		if l1 == nil && l2 == nil {
			if carry > 0 {
				lastNode.Next = &ListNode{Val: carry}
			}
			return firstNode
		}
		n := &ListNode{}
		v := carry
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		if v >= 10 {
			n.Val += v - 10
			carry = 1
		} else {
			n.Val += v
			carry = 0
		}

		if firstNode == nil {
			firstNode = n
		}
		if lastNode != nil {
			lastNode.Next = n
		}
		lastNode = n
	}
}