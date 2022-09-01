package main

import (
	"fmt"
	"math"
)

func main() {
	var length int
	fmt.Scan(&length)
	var arr = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}
	//double_pointer
	var j int = 0
	var max int
	var tempArr = make(map[int]int, length)
	for i := 0; i < length; i++ {
		tempArr[arr[i]]++
		for {
			value := tempArr[arr[i]]
			if value > 1 {
				tempArr[arr[j]]--
				j++
			} else {
				break
			}
		}
		max = int(math.Max(float64(i-j+1), float64(max)))
	}
	fmt.Print(max)
}
