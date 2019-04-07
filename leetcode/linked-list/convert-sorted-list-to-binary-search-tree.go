package main

type ListNode struct {
 	Val int
 	Next *ListNode
}
type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val:head.Val}
	}
	fast := head
	slow := head
	preSlow := head
	for fast != nil && fast.Next != nil{
		fast = fast.Next.Next
		preSlow = slow
		slow = slow.Next
	}
	preSlow.Next = nil
	return &TreeNode{
		Val:slow.Val,
		Left:sortedListToBST(head),
		Right:sortedListToBST(slow.Next),
	}
}

func main() {
	l := &ListNode{Val:-10, Next:&ListNode{Val:-3, Next:&ListNode{Val:0, Next:&ListNode{5, &ListNode{9, nil}}}}}
	sortedListToBST(l)
}