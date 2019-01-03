package leetcode

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

func printList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
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