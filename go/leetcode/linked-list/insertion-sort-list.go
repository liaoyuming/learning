package main


type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	target := head.Next
	tail := head
	nextHead := head.Next.Next
	for target != nil {
		if tail.Val > target.Val {
			tmpHead := head
			var pre *ListNode
			for tmpHead != tail.Next && target.Val > tmpHead.Val{
				pre = tmpHead
				tmpHead = tmpHead.Next
			}
			if tmpHead == head {
				head = target
			} else {
				pre.Next = target
			}
			target.Next = tmpHead
			tail.Next = nextHead
		}
		if tail.Next != target {
			target = nextHead
		} else {
			tail = tail.Next
			target = tail.Next
		}
		if nextHead != nil {
			nextHead = nextHead.Next
		}
	}
	return head
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func main() {
	PrintList(insertionSortList(&ListNode{1, &ListNode{15,&ListNode{2,&ListNode{10,nil}}}}))
}
