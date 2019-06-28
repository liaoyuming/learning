package main

func main() {
	//l1 := ListNode{Val: 1, Next:&(ListNode{Val: 2, Next:&(ListNode{Val: 1})})}
	//l2 := ListNode{Val: 9, Next:&(ListNode{Val: 9, Next:&(ListNode{Val: 1})})}
	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 9, Next:&(ListNode{Val: 9})}
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
	head := new(ListNode)
	res := head
	i := 0
	for l1 != nil || l2 != nil {
		temp := new(ListNode)
		l1Val := 0
		l2Val := 0
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		tempVal := l1Val + l2Val + i
		if tempVal >= 10 {
			temp.Val = tempVal - 10
			i = 1
		} else {
			temp.Val = tempVal
			i = 0
		}
		res.Next = temp
		res = res.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if i != 0 {
		res.Next = &(ListNode{Val:1})
	}
	return head.Next
}