package main

type MinStack struct {
	stack []int
	mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
	} else if this.GetMin() > x {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop()  {
	if this.stack[len(this.stack)-1] < this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {lengthOfLongestSubstring
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	println(minStack.GetMin())
	minStack.Pop()
	println(
		minStack.Top(),
		minStack.GetMin())
}