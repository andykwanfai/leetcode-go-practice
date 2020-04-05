package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	current := head
	n := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			n += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: n % 10, Next: nil}
		current = current.Next
		n /= 10
	}
	if n > 0 {
		current.Next = &ListNode{Val: n, Next: nil}
	}
	return head.Next
}
