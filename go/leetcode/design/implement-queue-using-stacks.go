package main

type MyQueue struct {
	inStack *[]int
	outStack *[]int
	inLength int
	outLength int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	inStack := []int{}
	outStack := []int{}
	return MyQueue{inStack:&inStack, outStack:&outStack}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	*(this.inStack) = append(*(this.inStack), x)
	this.inLength ++
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.outLength <= 0 {
		for i:=this.inLength-1; i>=0; i-- {
			*(this.outStack) = append(*(this.outStack), (*(this.inStack))[i])
		}
		this.inStack = &[]int{}
		this.outLength = this.inLength
		this.inLength = 0
	}
	res := (*(this.outStack))[this.outLength-1]
	*(this.outStack) = (*(this.outStack))[:this.outLength-1]
	this.outLength--
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.outLength <= 0 {
		for i:=this.inLength-1; i>=0; i-- {
			*(this.outStack) = append(*(this.outStack), (*(this.inStack))[i])
		}
		this.inStack = &[]int{}
		this.outLength = this.inLength
		this.inLength = 0
	}
	return (*(this.outStack))[this.outLength-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return (this.inLength + this.outLength) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	println(obj.Peek(), obj.Pop(),obj.Peek(),  obj.Empty())
}