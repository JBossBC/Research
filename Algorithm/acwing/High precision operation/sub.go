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
	// if Aarr>=Barr return true else return false
	if compareSize(Aarr, Barr) {
		sub(Aarr, Barr)
	} else {
		fmt.Printf("%s", "-")
		sub(Barr, Aarr)
	}
}
func compareSize(a []int, b []int) bool {
	if len(a) > len(b) {
		return true
	} else if len(a) < len(b) {
		return false
	} else {
		var length = len(a)
		for i := length - 1; i >= 0; i-- {
			if a[i] > b[i] {
				return true
			} else if a[i] < b[i] {
				return false
			}
		}
	}
	return true
}

// a>=b
//注意截尾
func sub(a []int, b []int) {
	var pointer = 0
	var maxLength = int(math.Max(float64(len(a)), float64(len(b))))
	var next = 0
	var result = make([]int, len(a))
	for pointer < maxLength {
		if pointer < len(b) {
			next -= b[pointer]
		}
		next += a[pointer]
		result[pointer] = next % 10
		if result[pointer] < 0 {
			next -= 10
			result[pointer] += 10
		}
		next = next / 10
		pointer++
	}
	var VaildIndex = len(a) - 1
	for VaildIndex > 0 && result[VaildIndex] == 0 {
		VaildIndex--
	}
	for i := VaildIndex; i >= 0; i-- {
		fmt.Printf("%d", result[i])
	}
}
