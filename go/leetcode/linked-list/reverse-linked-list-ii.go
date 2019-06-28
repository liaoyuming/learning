package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	res := &ListNode{}
	res.Next = head
	r := res
	for i:=0; i<m; i++ {
		r = r.Next
	}
	for i:=m; i<n; i++ {
		t := r.Next.Next
		r.Next.Next = t.Next
		t.Next = r.Next
		r.Next = t
	}
	return res.Next
}

func main() {
	head := ListNode{Val: 7, Next:&ListNode{Val:2, Next:&ListNode{Val:4, Next:&ListNode{Val:3}}}}
	l := reverseBetween(&head, 2, 3)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}
