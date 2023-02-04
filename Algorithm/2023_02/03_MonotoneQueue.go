package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var arrLength, windowsLength int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &arrLength, &windowsLength)
	var arr = make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Fscanf(reader, "%d", &arr[i])
	}
	var queue = make([]int, arrLength)
	var top, tail = 0, -1
	// min value
	for i := 0; i < len(arr); i++ {
		//如果已经超出了滑动窗口的大小,那么出队列
		if top <= tail && queue[top] < i-windowsLength+1 {
			top++
		}
		// 把控进来的元素
		for top <= tail && arr[queue[tail]] >= arr[i] {
			tail--
		}
		tail++
		queue[tail] = i
		if i >= windowsLength-1 {
			fmt.Printf("%d ", arr[queue[top]])
		}
	}
	fmt.Println()
	//reset params
	top, tail = 0, -1
	for i := 0; i < len(arr); i++ {
		if top <= tail && queue[top] < i-windowsLength+1 {
			top++
		}
		for top <= tail && arr[queue[tail]] <= arr[i] {
			tail--
		}
		tail++
		queue[tail] = i
		if i-windowsLength+1 >= 0 {
			fmt.Printf("%d ", arr[queue[top]])
		}
	}

}
