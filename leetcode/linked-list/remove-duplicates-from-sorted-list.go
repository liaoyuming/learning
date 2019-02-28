package main

/**
		删除排序链表中的重复元素
		https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
 */

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := head
	for head.Next != nil {
		if head.Next.Val == head.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}
	return res
}