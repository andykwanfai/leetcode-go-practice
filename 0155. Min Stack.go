package main

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(x int) {
	if len(this.min) == 0 || x <= this.GetMin() {
		this.min = append(this.min, x)
	}
	this.stack = append(this.stack, x)
}

func (this *MinStack) Pop() {
	stackLength := len(this.stack)
	if this.stack[stackLength-1] == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:stackLength-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
