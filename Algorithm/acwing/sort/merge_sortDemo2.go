package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var length int
	fmt.Fscan(reader, &length)
	var nums = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	fmt.Printf("%d", findReverseTimes(nums, 0, length-1))
}
func findReverseTimes(nums []int, l int, r int) int {
	if l >= r {
		return 0
	}
	var middle = (l + r) >> 1
	var left = findReverseTimes(nums, l, middle)
	var right = findReverseTimes(nums, middle+1, r)
	var mergeNumber = 0
	var i, j = l, middle + 1
	var tempArr = make([]int, r-l+1)
	var pointer = 0
	for i <= middle && j <= r {
		if nums[i] > nums[j] {
			tempArr[pointer] = nums[j]
			mergeNumber += middle + 1 - i
			pointer++
			j++
		} else {
			tempArr[pointer] = nums[i]
			pointer++
			i++
		}
	}
	for i <= middle {
		tempArr[pointer] = nums[i]

		pointer++
		i++
	}
	for j <= r {
		tempArr[pointer] = nums[j]
		pointer++
		j++

	}
	for z := 0; z < len(tempArr); z++ {
		nums[l+z] = tempArr[z]
	}
	return left + right + mergeNumber
}
