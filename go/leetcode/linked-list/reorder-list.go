package main

 type ListNode struct {
 	Val int
 	Next *ListNode
 }

func reorderList(head *ListNode)  {
	if head == nil || head.Next == nil {
		return
	}
	k := head
	s := head
	for k != nil && k.Next != nil {
		k = k.Next.Next
		s = s.Next
	}
	var revHead *ListNode
	revHead.Next = nil
	for s != nil && s.Next != nil {
		t := s.Next
		s.Next = s.Next.Next
		t.Next = revHead
		revHead = t
	}
	for head != nil && revHead != nil {
		t := head.Next
		r := revHead.Next
		head.Next = revHead
		revHead.Next = t
		head = t
		revHead = r
	}
}