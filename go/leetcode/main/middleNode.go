package main

/**
		链表的中间结点
		https://leetcode-cn.com/problems/middle-of-the-linked-list/description/
 */

func main() {
	head := &ListNode{Val: 2, Next:&ListNode{Val:2, Next:&ListNode{Val:2}}}
	//head := &ListNode{}
	l := middleNode(head)
	printList(l)
}
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

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	q := head
	slow := head
	for slow != nil && slow.Next != nil {
		q = q.Next
		slow = slow.Next.Next
	}
	return q
}