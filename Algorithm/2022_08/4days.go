package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{1, 2}
	println("arr: ", cap(arr))
	println("arr: ", len(arr))
	var arr1 = append(arr, 1, 2, 3, 4, 5, 6, 7)
	println("arr1: ", cap(arr1))
	println("arr1: ", len(arr1))
	fmt.Println(arr1)
}

func minSubsequence(nums []int) []int {
	var arr2 = []int{1, 2, 3, 4, 5, 6, 7}
	println(cap(arr2))
	println(len(arr2))
	var arr3 = append(arr2, 1)
	println(cap(arr3))
	println(len(arr3))
	//贪婪算法
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			return false
		}
		return true
	})
	fmt.Println(nums)
	var result = make([]int, 0)
	if nums == nil || len(nums) == 0 {
		return result
	}
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	var leftSum = 0
	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		sum -= nums[i]
		result = append(result, nums[i])
		if leftSum > sum {
			return result
		}
	}
	return nil
}
