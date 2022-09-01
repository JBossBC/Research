package main

import (
	"fmt"
	"sort"
)

func main() {
	var result = findClosestElements([]int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}, 3, 5)
	for _, value := range result {
		fmt.Printf("%d", value)
	}
}

//双指针问题
func findClosestElements(arr []int, k int, x int) []int {
	if arr[0] >= x {
		return arr[:k]
	}
	if arr[len(arr)-1] <= x {
		return arr[len(arr)-k:]
	}
	var a, b = 0, len(arr) - 1
	for a < b {
		var middle = (a + b) >> 1
		if arr[middle] >= x {
			b = middle
		} else {
			a = middle + 1
		}
	}
	var l, r = a - 1, a
	var result = make([]int, k)
	var pointer = 0
	for k > 0 && r < len(arr) && l >= 0 {
		if arr[r]-x < x-arr[l] {
			result[pointer] = arr[r]
			pointer++
			r++
		} else {
			result[pointer] = arr[l]
			pointer++
			l--
		}
		k--
	}
	if k > 0 {
		for l >= 0 && k > 0 {
			result[pointer] = arr[l]
			pointer++
			l--
			k--
		}
		for r < len(arr) && k > 0 {
			result[pointer] = arr[r]
			pointer++
			r++
			k--
		}
	}
	sort.Ints(result)
	return result
}
