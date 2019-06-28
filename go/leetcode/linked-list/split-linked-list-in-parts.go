package main

type ListNode struct {
	Val int
 	Next *ListNode
 }

func splitListToParts(root *ListNode, k int) []*ListNode {
	res := []*ListNode{}
	if k == 0 {
		return res
	}
	length := 0
	t := root
	for t != nil {
		length++
		t = t.Next
	}
	m := length / k
	n := length % k
	for j:=0; j<k; j++ {
		temp := &ListNode{}
		tempHead := temp
		target := m
		if n > 0 {
			n--
			target++
		}
		for i:=0; i<target && root!=nil; i++ {
			temp.Next = root
			root = root.Next
			temp = temp.Next
		}
		temp.Next = nil
		res = append(res, tempHead.Next)
	}
	return res
}

func PrintList(l *ListNode)  {
	for l != nil {
		println(l.Val)
		l = l.Next
	}
	println("###")
}

func PrintLists(l []*ListNode)  {
	for i:=0; i<len(l); i++ {
		PrintList(l[i])
	}
}


func main() {
	PrintLists(splitListToParts(&ListNode{1,&ListNode{4,&ListNode{5,nil}}}, 2))
}