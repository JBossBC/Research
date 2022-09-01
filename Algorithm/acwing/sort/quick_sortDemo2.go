package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var length int
	var location int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &length)
	fmt.Fscan(reader, &location)
	var nums = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &nums[i])
	}
	quickSort(nums, 0, length-1)
	print(nums[location-1])
}

//TODO
func quickSort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	var middle, i, j = nums[l], l - 1, r + 1
	for i < j {
		for {
			i++
			if nums[i] >= middle {
				break
			}
		}
		for {
			j--
			if nums[j] <= middle {
				break
			}
		}
		if i < j {
			nums[i] = nums[i] ^ nums[j]
			nums[j] = nums[i] ^ nums[j]
			nums[i] = nums[i] ^ nums[j]
		}
	}
	quickSort(nums, l, j)
	quickSort(nums, j+1, r)
}
