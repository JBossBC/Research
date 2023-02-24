package main

import "math"

func main() {

}

func minimumOperations(nums []int) int {
	var s []bool = make([]bool, len(nums))
	var n = len(nums)
	var times = 0
	for n != 0 {
		var min int = math.MaxInt
		for i := 0; i < len(nums); i++ {
			if nums[i] != 0 && nums[i] < min {
				min = nums[i]
			}
		}
		if min == math.MaxInt {
			return times
		}
		for i := 0; i < len(nums); i++ {
			if nums[i]-min <= 0 {
				nums[i] = 0
			} else {
				nums[i] = nums[i] - min
			}
			if !s[i] && nums[i] == 0 {
				s[i] = true
				n--
			}
		}
		times++
	}
	return times
}
