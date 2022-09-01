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
	mergeSort(nums, 0, length-1)
	for _, value := range nums {
		fmt.Printf("%d ", value)
	}
}
func mergeSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	var middle = (l + r) >> 1
	mergeSort(nums, l, middle)
	mergeSort(nums, middle+1, r)
	var temp = make([]int, r-l+1)
	var i, j = l, middle + 1
	var pointer = 0
	for i <= middle && j <= r {
		if nums[i] > nums[j] {
			temp[pointer] = nums[j]
			pointer++
			j++
		} else {
			temp[pointer] = nums[i]
			pointer++
			i++
		}
	}
	for i <= middle {
		temp[pointer] = nums[i]
		i++
		pointer++
	}
	for j < r {
		temp[pointer] = nums[j]
		j++
		pointer++
	}
	for k := 0; k < len(temp); k++ {
		nums[l+k] = temp[k]
	}
}
