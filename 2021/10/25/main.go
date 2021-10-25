package main

type MinStack struct {
	stack [][]int
}

func Constructor() MinStack {
	return MinStack{
		stack: [][]int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, []int{val, val})
	} else {
		this.stack = append(this.stack, []int{val, this.Min(val)})
	}
}

func (this *MinStack) Min(val int) int {
	if val < this.GetMin() {
		return val
	}
	return this.GetMin()
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
