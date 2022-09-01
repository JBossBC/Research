package main

import (
	"fmt"
)

func main() {
	var A string
	var B int
	fmt.Scanf("%s %d", &A, &B)
	var aArr = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		aArr[i] = int(A[i] - '0')
	}
	mul(aArr, B)
}

func mul(A []int, B int) {
	var result = make([]int, len(A)+B)
	var next = 0
	var pointer = 0
	for {
		if len(A) > pointer {
			next += A[pointer] * B
		}
		result[pointer] = next % 10
		next = next / 10
		pointer++
		if next == 0 && pointer >= len(A) {
			break
		}
	}
	var guessLocation = len(result) - 1
	for guessLocation > 0 && result[guessLocation] == 0 {
		guessLocation--
	}
	for i := guessLocation; i >= 0; i-- {
		fmt.Printf("%d", result[i])
	}
}
