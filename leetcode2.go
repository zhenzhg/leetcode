package main

//import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{-1, nil}
	carry := 0
	p, q, curr := l1, l2, dummyHead

	for p != nil || q != nil {
		x := 0
		if p != nil {
			x = p.Val
		}
		y := 0
		if q != nil {
			y = q.Val
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry, Next: nil}
	}
	return dummyHead.Next
}
