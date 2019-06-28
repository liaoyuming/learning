package main

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left := &ListNode{}
	right := &ListNode{}
	leftHead := left
	rightHead := right
	for head != nil {
		if head.Val < x {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	right.Next = nil
	left.Next = rightHead.Next
	return leftHead.Next
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func main() {
	PrintList(partition(&ListNode{5, &ListNode{2,&ListNode{1,nil}}}, 5))
}
