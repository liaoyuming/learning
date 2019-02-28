package main

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	length := 0
	t := head
	for t != nil {
		length ++
		t = t.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	n := length - k
	resHead := head
	for i:=n; i > 0; i-- {
		if i == 1 {
			next := resHead.Next
			resHead.Next = nil
			resHead = next
		} else {
			resHead = resHead.Next
		}
	}
	temp := resHead
	for temp != nil && temp.Next != nil {
		temp = temp.Next
	}
	if temp != nil {
		temp.Next = head
	}
	return resHead
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func main() {
	PrintList(rotateRight(&ListNode{1,&ListNode{4,&ListNode{5,nil}}}, 3))
}