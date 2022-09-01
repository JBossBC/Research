package main

import (
	"fmt"
	"math"
)

/* challenge failed*/
func main() {
	//println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	//println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}))
	for i := 0; i < 100; i++ {

		fmt.Printf("%d     ", maxEqualFreq([]int{1, 2, 3, 1, 2, 3, 4, 4, 4, 4, 1, 2, 3, 5, 6}))
	}
}

func maxEqualFreq(nums []int) int {
	var originValue = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		originValue[nums[i]]++
	}
	var actualValue = -1
	for i := len(nums) - 1; i >= 0; i-- {
		var mapValue = originValue
		var deleteChance = false
		var success = true
		var hasBothEqual = false
		var maxValueTimes = math.MaxInt32
		var occurTimes = make(map[int]int, len(nums))
		for _, value := range mapValue {
			occurTimes[value]++
		}
		var occurTempValue = -1
		var hasEqualSize = false
		for index, value := range occurTimes {
			if value > occurTempValue {
				maxValueTimes = index
				occurTempValue = value
				hasEqualSize = false
			}
			if value == occurTempValue {
				hasEqualSize = true
			}
		}
		if hasEqualSize {

		}
		actualValue = maxValueTimes
		for index, value := range mapValue {
			if value == 0 {
				delete(mapValue, index)
				continue
			}
			if value == actualValue {
				hasBothEqual = true
				continue
			}
			if value != actualValue {
				if deleteChance {
					success = false
					break
				}
				if hasBothEqual && value == actualValue+1 {
					deleteChance = true
					continue
				}
				if value == 1 {
					deleteChance = true
					delete(mapValue, index)
					continue
				}
				if !hasBothEqual && value == actualValue-1 {
					deleteChance = true
					hasBothEqual = true
					actualValue = value
				} else {
					success = false
				}
			}
			if !success {
				break
			}

		}
		if !deleteChance {
			success = false
		}
		if success {
			return i + 1
		}
		tempValue = -1
		actualValue = -1
		mapValue[nums[i]]--
	}
	return -1
}
