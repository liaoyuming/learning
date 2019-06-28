package main

/**
 		两数相加 II
		https://leetcode-cn.com/problems/add-two-numbers-ii/description/
 */

func main() {
	l1 := ListNode{Val: 7, Next:&ListNode{Val:2, Next:&ListNode{Val:4, Next:&ListNode{Val:3}}}}
	l2 := ListNode{Val: 5, Next:&(ListNode{Val: 6, Next:&ListNode{Val:4}})}
	l := addTwoNumbers(&l1, &l2)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var list1 *ListNode
	var list2 *ListNode
	n := 0
	for l1 != nil  {
		temp := ListNode{}
		temp = *l1
		temp.Next = list1
		list1 = &temp
		l1 = l1.Next
	}
	for l2 != nil  {
		temp := ListNode{}
		temp = *l2
		temp.Next = list2
		list2 = &temp
		l2 = l2.Next
	}
	for list1 != nil || list2 != nil {
		val1 := 0
		val2 := 0
		if list1 != nil {
			val1 = list1.Val
		}
		if list2 != nil {
			val2 = list2.Val
		}
		n = val1 + val2 + n
		r := &ListNode{Val:n%10}
		r.Next = res
		res = r
		n = n/10
		if list1 != nil {
			list1 = list1.Next
		}
		if list2 != nil {
			list2 = list2.Next
		}
	}
	if n == 1 {
		r := &ListNode{Val:1}
		r.Next = res
		res = r
	}
	return res
}