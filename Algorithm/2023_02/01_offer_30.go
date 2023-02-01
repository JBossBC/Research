package main

import (
	"math"
	"sync/atomic"
)

type MinStack struct {
	data     []int
	top      int64
	minValue []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var currentValue = MinStack{
		data:     make([]int, 2),
		top:      0,
		minValue: make([]int, 2),
	}
	currentValue.data[0] = math.MaxInt64

	return currentValue
}

func (this *MinStack) Push(x int) {
	atomic.AddInt64(&this.top, 1)
	if int64(len(this.data)) <= this.top {
		this.data = append(this.data, x)
	} else {
		this.data[this.top] = x
	}
	var preMin = this.data[this.minValue[this.top-1]]
	var currentValue = x
	var MinValue = -1
	if currentValue < preMin {
		MinValue = int(this.top)
	} else {
		MinValue = this.minValue[this.top-1]
	}

	if int64(len(this.minValue)) <= this.top {
		this.minValue = append(this.minValue, MinValue)
	} else {
		this.minValue[this.top] = MinValue
	}
}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
	}
	atomic.AddInt64(&this.top, -1)
}

func (this *MinStack) Top() int {
	if this.top <= 0 {
		return math.MinInt64
	}
	return this.data[this.top]
}

func (this *MinStack) Min() int {
	return this.data[this.minValue[this.top]]
}

func main() {
	stack := Constructor()
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
}
