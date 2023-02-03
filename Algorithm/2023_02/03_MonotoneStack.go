package main

import (
	"fmt"
	"math"
)

func main() {
	var length int
	fmt.Scanln(&length)
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	var stack = make([]int, length)
	var topPointer = 0
	var result = make([]int, length)
	stack[topPointer] = math.MaxInt
	for i := 0; i < length; i++ {
		var find bool
		var tempPointer = topPointer
		for tempPointer != 0 {
			if stack[tempPointer] < arr[i] {
				find = true
				result[i] = stack[tempPointer]
				break
			}
			tempPointer--
		}
		if !find {
			result[i] = -1
		}
		// can join monotoneStack
		if arr[i] <= stack[topPointer] {
			topPointer++
			stack[topPointer] = arr[i]
		}

	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", result[i])
	}

}
