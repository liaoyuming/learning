package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	h := head
	for h != nil {
		h = h.Next
		count++
		if count == k {
			var reverseHead = reverseKGroup(h, k)
			for count > 0 {
				node := head
				head = head.Next
				node.Next = reverseHead
				reverseHead = node
				count--
			}
			return reverseHead
		}
	}
	return head
}

func main() {
	head := &ListNode{Val: 1, Next:&ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:4, Next:&ListNode{Val:5}}}}}
	l := reverseKGroup(head, 2)
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}