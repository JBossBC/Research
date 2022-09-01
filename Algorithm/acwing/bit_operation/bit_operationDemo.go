package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scan(&length)
	var numberArr = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&numberArr[i])
	}
	var resultArr = make([]int, length)
	var pointer = 0
	//补码等于反码+1=>x+(-x)=0
	for i := 0; i < length; i++ {
		var value = numberArr[i]
		var result = 0
		for value != 0 {
			var temp = (-value) & value
			if temp != 0 {
				result++
			}
			value -= temp
		}
		resultArr[pointer] = result
		pointer++
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", resultArr[i])
	}
}
