package main

import (
	"fmt"
	"math"
)

func main() {
	var arrLength, windowsLength int
	fmt.Scanln(&arrLength, &windowsLength)
	var arr = make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var queue = make([]int, arrLength+1)
	queue[0] = math.MaxInt
	var top, tail = 0, 1
	// min value
	for i := 0; i < arrLength; i++ {
		if top <= tail && i-windowsLength+1 > queue[top] {
			top++
		}
		for top <= tail && arr[queue[tail]] >= arr[i] {
			tail--
		}
		tail++
		queue[tail] = i
		if i-windowsLength-1 >= 0 {
			fmt.Printf("%d", arr[queue[top]])
		}

	}

}
