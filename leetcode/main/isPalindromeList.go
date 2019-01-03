package leetcode

/**
		回文链表
		https://leetcode-cn.com/problems/palindrome-linked-list/description/
 */

func main() {
	//head := &ListNode{Val: 2, Next:&ListNode{Val:2, Next:&ListNode{Val:2, Next:&ListNode{Val:1}}}}
	head := &ListNode{}
	b := isPalindrome(head)
	println(b)
}

 type ListNode struct {
 	Val int
 	Next *ListNode
 }


func printList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func isPalindrome(head *ListNode) bool {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	len := len(list)
	for i:=0; i<len/2+1; i++ {
		if list[i] != list[len-1-i] {
			return false
		}
	}
	return true
}