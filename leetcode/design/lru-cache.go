package main

type ListNode struct {
	Val int
	Key int
	Next *ListNode
	Pre *ListNode
}

type LRUCache struct {
	head *ListNode
	tail *ListNode
	hash map[int]*ListNode
	capacity int
	len int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity:capacity, len:0, hash:map[int]*ListNode{}}
}

func (this *LRUCache) Get(key int) int {
	if l, ok := this.hash[key]; ok {
		if l.Pre != nil {
			l.Pre.Next = l.Next
			if l.Next != nil {
				l.Next.Pre = l.Pre
			}
			if this.tail == l {
				this.tail = l.Pre
			}
			l.Next = this.head
			l.Pre = nil
			this.head.Pre = l
			this.head = l
		}
		return l.Val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	if l, ok := this.hash[key]; ok {
		if l.Pre != nil {
			l.Pre.Next = l.Next
			if l.Next != nil {
				l.Next.Pre = l.Pre
			}
			if this.tail == l {
				this.tail = l.Pre
			}
			l.Next = this.head
			l.Pre = nil
			this.head.Pre = l
			this.head = l
		}
		l.Val = value
	} else {
		l := &ListNode{Val:value, Key:key}
		l.Next = this.head
		if this.head != nil {
			this.head.Pre = l
		}
		this.head = l
		this.len++
		if this.len == 1 {
			this.tail = l
		} else if this.len > this.capacity {
			delete(this.hash, this.tail.Key)
			this.tail = this.tail.Pre
		}
		this.hash[key] = l
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(1,1)
	obj.Put(2,2)
	obj.Get(1)
	obj.Put(3,3)
	//obj.Put(2,2)
	println(obj.Get(2))
}


