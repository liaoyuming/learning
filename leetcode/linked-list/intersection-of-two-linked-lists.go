package main

type ListNode struct {
	Val int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tA := headA
	countA := 0
	for tA != nil {
		countA ++
		tA = tA.Next
	}
	tB := headB
	countB := 0
	for tB != nil {
		countB ++
		tB = tB.Next
	}
	n := 0
	if countA > countB {
		n = countA - countB
		for i:=0; i < n; i++ {
			headA = headA.Next
		}
	} else {
		n = countB - countA
		for i:=0; i < n; i++ {
			headB = headB.Next
		}
	}
	for headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	print(getIntersectionNode2(&ListNode{1,nil }, &ListNode{2, &ListNode{1,nil }}))

}