package main

import "fmt"

func main() {
	arr := sortArray([]int{5, 1, 1, 2, 0, 0})
	fmt.Println(arr)
}

//func sortArray(nums []int) []int {
//	merge_sort(nums, 0, len(nums)-1)
//	return nums
//}

/**
  归并排序的核心思想是减治法，对于减治法，很重要的是拆分问题的规模，通过逐步解决小规模问题达到解决大规模问题的理念，考虑这类问题的时候
  应该以抽象的思维去思考,不能以具体的规模数组去企图寻找解决问题的方法，这会限制你对整体小规模问题通解的把握。
  归并排序的模板:
  ①、确定分界点
  ②、递归排序(注意死循环,边界的模糊容易造成死循环,middle取值的向下取整如果考虑不清楚极容易出现死循环)
  ③、归并!!!(注意每个元素都要被排序，不要遗漏)
*/
func merge_sort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	var middle = (l + r) >> 1
	merge_sort(nums, l, middle)
	merge_sort(nums, middle+1, r)
	var temp = make([]int, r-l+1)
	var tempIndex = 0
	var mid = middle + 1
	var i = l
	for i <= middle && mid <= r {
		if nums[i] > nums[mid] {
			temp[tempIndex] = nums[mid]
			mid++
		} else {
			temp[tempIndex] = nums[i]
			i++
		}
		tempIndex++
	}
	for i <= middle {
		temp[tempIndex] = nums[i]
		i++
		tempIndex++
	}
	for mid <= r {
		temp[tempIndex] = nums[mid]
		mid++
		tempIndex++
	}
	var j = l
	for k := 0; k < len(temp); k++ {
		nums[j] = temp[k]
		j++
	}
}
