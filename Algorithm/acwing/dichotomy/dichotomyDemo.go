package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var systemIn = bufio.NewReader(os.Stdin)
	var length int
	fmt.Fscan(systemIn, &length)
	var times int
	fmt.Fscan(systemIn, &times)
	var nums = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(systemIn, &nums[i])
	}
	var needFind = make([]int, times)
	var pointer = 0
	for i := 0; i < times; i++ {
		fmt.Fscan(systemIn, &needFind[pointer])
		pointer++
	}
	for _, value := range needFind {
		dichotomy_Search(nums, value)
	}

}

/**
  二分不仅仅适用于有序的情况，只需要有特定的分解条件即可
  两个模板:当r=middle 时，初始的middle=(l+r)>>1;当 r=middle-1时,初始的middle=(l+r+1)>>1;
  这两个模板对应的情况为:求分界点左边;求分界点右边
*/
func dichotomy_Search(nums []int, target int) {
	var l, r = 0, len(nums) - 1
	for l < r {
		var middle = (l + r) >> 1
		if nums[middle] >= target {
			r = middle
		} else {
			l = middle + 1
		}
	}
	if nums[l] != target {
		fmt.Printf("%d %d", -1, -1)
		fmt.Println()
		return
	} else {
		var k = len(nums) - 1
		for l < k {
			var middle = (l + k + 1) >> 1
			if nums[middle] <= target {
				l = middle
			} else {
				k = middle - 1
			}
		}
	}
	fmt.Printf("%d %d", r, l)
	fmt.Println()
	return
}
