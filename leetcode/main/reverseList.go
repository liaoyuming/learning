package leetcode
/**
	反转一个单链表
	leetcode : https://leetcode-cn.com/problems/reverse-linked-list/description/
 */
type ListNode struct {
	Val int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	var reverseHead *ListNode

	for ; head != nil; {
		next := head.Next
		head.Next = reverseHead
		reverseHead = head
		head = next
	}
	return reverseHead
}

// @todo  使用递归
func reverseList2(head *ListNode) *ListNode {

}

func main()  {
	var head *ListNode
	arr := [5]int{5, 4, 3, 2, 1}
	for i:=0; i<len(arr); i++ {
		var tmp = new(ListNode)
		tmp.Val = arr[i]
		tmp.Next = head
		head = tmp

	}
	//for ; head != nil; {
	//	println(head.Val)
	//	head = head.Next
	//}

	reverseHead := reverseList(head)

	for ; reverseHead != nil; {
		println(reverseHead.Val)
		reverseHead = reverseHead.Next
	}
}