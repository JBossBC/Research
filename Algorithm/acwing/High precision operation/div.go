package main

import (
	"fmt"
)

func main() {
	var A string
	var B int
	fmt.Scan(&A)
	fmt.Scan(&B)
	var aArr = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		aArr[i] = int(A[i] - '0')
	}
	div(aArr, B)
}

// a不需要换顺序
func div(a []int, b int) {
	var result = make([]int, len(a))
	var next = 0
	var pointer = 0
	for pointer < len(a) {
		next = next*10 + a[pointer]
		result[pointer] = next / b
		next %= b
		pointer++
	}

	var ValidIndex int
	for ValidIndex < len(result)-1 && result[ValidIndex] == 0 {
		ValidIndex++
	}
	for ValidIndex < len(result) {
		fmt.Printf("%d", result[ValidIndex])
		ValidIndex++
	}
	fmt.Println()
	fmt.Printf("%d", next)
}
