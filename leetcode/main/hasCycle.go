package main

/**
		环形链表
		https://leetcode-cn.com/problems/linked-list-cycle/description/
 */

func main() {
	//head := &ListNode{Val: 2, Next:&ListNode{Val:2, Next:&ListNode{Val:2}}}
	head := &ListNode{}
	b := hasCycle(head)
	println(b)
}

type ListNode struct {
	Val int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	s := head
	q := head
	for q != nil && q.Next != nil {
		s = s.Next
		q = q.Next.Next
		if s == q {
			return true
		}
	}
	return false
}