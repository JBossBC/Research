package main

import (
	"fmt"
	"math"
)

func main() {
	var a string
	var b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	var Aarr = make([]int, len(a))
	var Barr = make([]int, len(b))
	for i := 0; i < len(Aarr); i++ {
		Aarr[i] = int(a[len(a)-i-1] - '0')
	}
	for i := 0; i < len(Barr); i++ {
		Barr[i] = int(b[len(b)-i-1] - '0')
	}
	add(Aarr, Barr)
}

func add(a []int, b []int) {
	var next = 0
	var pointer = 0
	var maxLength = int(math.Max(float64(len(a)), float64(len(b))))
	var result = make([]int, maxLength+1)
	for pointer <= maxLength {
		if pointer < len(a) {
			next += a[pointer]
		}
		if pointer < len(b) {
			next += b[pointer]
		}
		result[pointer] = next % 10
		next = next / 10
		pointer++
	}
	var ValidIndex = len(result) - 1
	for ValidIndex >= 0 {
		if result[ValidIndex] != 0 {
			break
		}
		ValidIndex--
	}
	for i := ValidIndex; i >=0; i-- {
		fmt.Printf("%d", result[i])
	}
}
