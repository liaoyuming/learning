package main

/**
		合并两个有序链表
		https://leetcode-cn.com/problems/merge-two-sorted-lists/
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 4}
	l2 := ListNode{Val: 2, Next:&(ListNode{Val: 9})}
	l := mergeTwoLists(&l1, &l2)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	res := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Next = l1
			l1 = l1.Next
		} else {
			res.Next = l2
			l2 = l2.Next
		}
		res = res.Next
	}
	if l1 != nil {
		res.Next = l1
	} else if l2 != nil {
		res.Next = l2
	}
	return head.Next
}