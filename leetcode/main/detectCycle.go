package main

/**
		环形链表 II
		https://leetcode-cn.com/problems/linked-list-cycle-ii/description/
 */

func main() {
	a := &ListNode{Val:2}
	b := &ListNode{Val:3}
	a.Next = b
	b.Next = a
	//head := &ListNode{Val: 2, Next:&ListNode{Val:2, Next:&ListNode{Val:2}}}
	//head := &ListNode{}
	l := detectCycle(a)
	printList(l)
}

func printList(l *ListNode)  {
	n := 0
	for l != nil && n < 3 {
		println(l.Val)
		l = l.Next
		n ++
	}
	println("###")
}

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	b := false
	s := head
	fast := head
	for fast != nil && fast.Next != nil {
		s = s.Next
		fast = fast.Next.Next
		if s == fast {
			b = true
			break
		}
	}
	if !b {
		return nil
	}
	start := head
	for start != fast {
		start = start.Next
		fast = fast.Next
	}
	return start
}
