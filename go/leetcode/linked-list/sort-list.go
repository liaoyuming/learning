package main

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList2(head *ListNode) *ListNode {
	return quickSortList(head, nil)
}

func quickSortList(head *ListNode, tail *ListNode) *ListNode {
	if head == tail || head.Next == tail {
		return head
	}
	left := &ListNode{}
	leftHead := left
	right := &ListNode{}
	rightHead := right
	mid := head
	head = head.Next
	for head != tail {
		if head.Val < mid.Val {
			left.Next = head
			left = left.Next
		} else {
			right.Next = head
			right = right.Next
		}
		head = head.Next
	}
	left.Next = mid
	right.Next = tail
	leftHead = quickSortList(leftHead.Next, mid)
	rightHead = quickSortList(rightHead.Next, tail)
	mid.Next = rightHead
	return leftHead
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	quick := head
	slow := head
	preSlow := head
	for quick != nil && quick.Next != nil {
		preSlow = slow
		quick = quick.Next.Next
		slow = slow.Next
	}
	preSlow.Next = nil
	left := sortList(head)
	right := sortList(slow)
	res := &ListNode{}
	resHead := res
	for left != nil && right != nil {
		if left.Val < right.Val {
			res.Next = left
			left = left.Next
		} else {
			res.Next = right
			right = right.Next
		}
		res = res.Next
	}
	if left == nil {
		res.Next = right
	} else if right == nil {
		res.Next = left
	}
	return resHead.Next
}


func printList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("####")
}


func main() {
	head := &ListNode{Val:-1, Next:&ListNode{Val:5, Next:&ListNode{Val:3, Next:&ListNode{Val:4, Next:&ListNode{Val:0}}}}}
	//head := &ListNode{Val:4, Next:&ListNode{Val:2}}
	head = sortList(head)
	printList(head)
}