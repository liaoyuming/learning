package leetcode

/**
		反转链表 II
		https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 */

func main() {
	head := ListNode{Val: 7, Next:&ListNode{Val:2, Next:&ListNode{Val:4, Next:&ListNode{Val:3}}}}
	l := reverseBetween(&head, 2, 3)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
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
// @todo  需要再理解一波
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	res := &ListNode{}
	res.Next = head
	c := res
	for i:=1; i<m; i++ {
		c = c.Next
	}
	for i:=m; i<n; i++ {
		t := c.Next.Next
		c.Next.Next = t.Next
		t.Next = c.Next
		c.Next = t
	}
	return res.Next
}