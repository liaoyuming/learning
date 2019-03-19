package main

type MyStack struct {
	myList []int
	length int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	newList := []int{x}
	for i:=0; i<this.length; i++ {
		newList = append(newList, this.myList[i])
	}
	this.myList = newList
	this.length++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	res := this.myList[0]
	this.myList = this.myList[1:]
	this.length--
	return res
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.myList[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return  this.length <= 0
}

func main() {
	obj := Constructor()
	obj.Push(1)
	println(obj.Top(), obj.Pop())
	obj.Push(2)
	println(obj.Empty())
}
/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

