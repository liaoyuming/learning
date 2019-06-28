package main

/**
		删除排序链表中的重复元素 II
		https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
 */

func main() {
	head := &ListNode{Val: 1, Next:&ListNode{Val:2, Next:&ListNode{Val:3}}}
	//head := &ListNode{}
	l := deleteDuplicates(head)
	printList(l)
}

type ListNode struct {
	Val int
	Next *ListNode
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	list := &ListNode{}
	list.Next = head
	res := list
	for list.Next != nil && list.Next.Next != nil {
		if list.Next.Val == list.Next.Next.Val {
			tmp := list.Next.Val
			for list.Next != nil && list.Next.Val == tmp {
				list.Next = list.Next.Next
			}
		} else {
			list = list.Next
		}
	}
	return res.Next
}