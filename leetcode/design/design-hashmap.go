package main

type ListNode struct {
	Key int
	Val int
	Next *ListNode
}

type MyHashMap struct {
	data []*ListNode
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	data := make([]*ListNode, 500)
	return MyHashMap{data}
}

func (this *MyHashMap) hashKey(key int) int {
	return key % 500
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
	hashKey := this.hashKey(key)
	head := this.data[hashKey]
	for head != nil {
		if head.Key == key {
			head.Val = value
			return
		}
	}
	node := &ListNode{Key:key, Val:value}
	node.Next = head
	this.data[hashKey] = node
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashKey := this.hashKey(key)
	head := this.data[hashKey]
	for head != nil {
		if head.Key == key {
			return head.Val
		}
	}
	return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
	hashKey := this.hashKey(key)
	head := this.data[hashKey]
	var preNode, nextNode *ListNode
	for head != nil {
		if head.Key == key {
			nextNode = head.Next
			break
		}
		preNode = head
	}
	if preNode == nil {
		this.data[hashKey] = nextNode
	} else {
		preNode.Next = nextNode
	}
}
