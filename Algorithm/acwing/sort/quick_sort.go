package main

import (
	"bufio"
	"fmt"
	"os"
)

//func main() {
//	arr := sortArray([]int{5, 1, 1, 2, 0, 0})
//	fmt.Println(arr)
//}
//func sortArray(nums []int) []int {
//	quick_sort(nums, 0, len(nums)-1)
//	return nums
//}

/**
快速排序的模板:
  ①、寻找分界点
  ②、排序(快速排序的思想是将数组以分界点为标准划分为准两个数组,左边的数组一定小于等于分界点,右边的数组一定大于等于分界点,能够实现的方法很多,但核心思想是这样)
  ③、递归
*/
// func quick_sort(nums []int, l int, r int) {
// 	if l >= r {
// 		return
// 	}
// 	var middle = nums[l]
// 	var i = l - 1
// 	var j = r + 1
// 	for i < j {
// 		i++
// 		j--
// 		for nums[i] < middle {
// 			i++
// 		}

// 		for nums[j] > middle {
// 			j--
// 		}
// 		if i < j {
// 			nums[i] = nums[i] ^ nums[j]
// 			nums[j] = nums[i] ^ nums[j]
// 			nums[i] = nums[i] ^ nums[j]
// 		}
// 	}
// 	quick_sort(nums, l, j)
// 	quick_sort(nums, j+1, r)
// }

func main() {
	var length int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &length)
	q := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &q[i])
	}
	quick_sort(q, 0, length-1)
	for i := 0; i < length; i++ {
		fmt.Printf("%d ", q[i])
	}
}

func quick_sort(nums []int, l int, r int) {
	if l >= r {
		return
	}
	var middle = nums[l]
	var j, k = l - 1, r + 1
	for j < k {
		j++
		k--
		for nums[j] < middle {
			j++
		}
		for nums[k] > middle {
			k--
		}
		if j < k {
			nums[j] = nums[j] ^ nums[k]
			nums[k] = nums[j] ^ nums[k]
			nums[j] = nums[j] ^ nums[k]
		}

	}
	quick_sort(nums, l, k)
	quick_sort(nums, k+1, r)
}
