package main

import (
	"fmt"
)

func main() {
	fmt.Printf("this element %t exist in this arr", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
}
func searchMatrix(matrix [][]int, target int) bool {
	var l, r = 0, len(matrix) - 1
	for l < r {
		var middle = (l + r) >> 1
		if matrix[middle][0] >= target {
			r = middle
		} else {
			l = middle + 1
		}
	}
	if l > 0 && matrix[l][0] > target {
		l--
	}
	var x, y = 0, len(matrix[0]) - 1
	for x < y {
		var middle = (x + y) >> 1
		if matrix[l][middle] >= target {
			y = middle
		} else {
			x = middle + 1
		}
	}
	if matrix[l][x] != target {
		return false
	} else {
		return true
	}
}
