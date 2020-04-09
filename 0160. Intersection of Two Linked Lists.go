package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ptA := headA
	ptB := headB
	for ptA != ptB {
		if ptA == nil {
			ptA = headB
		} else {
			ptA = ptA.Next
		}
		if ptB == nil {
			ptB = headA
		} else {
			ptB = ptB.Next
		}
	}
	return ptA
}
