package main

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 1 {
		return lists[0]
	} else if length == 0 {
		return nil
	}
	left := mergeKLists(lists[:length/2])
	right := mergeKLists(lists[length/2:])
	res := &ListNode{}
	resHead := res
	for left != nil && right != nil{
		if left.Val < right.Val {
			res.Next = left
			left = left.Next
		} else {
			res.Next = right
			right = right.Next
		}
		res = res.Next
	}
	if left == nil && right != nil {
		res.Next = right
	} else if left != nil && right == nil {
		res.Next = left
	}
	return resHead.Next
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func main() {
	lists := []*ListNode{
		&ListNode{1,&ListNode{4,&ListNode{5,nil}}},
		&ListNode{1,&ListNode{3,&ListNode{4,nil}}},
		&ListNode{2,&ListNode{6, nil}},
	}
	PrintList(mergeKLists(lists))
}