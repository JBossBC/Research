package main

import (
	"fmt"
)

//切片复制的时候，底层数组不变，只是指针会有所变化
func main() {
	var arr = []int{1, 2, 3, 4, 5}
	var arr1 = arr[:1]
	fmt.Println(arr1[0])
	// fmt.Println(arr1[1])
	arr1 = append(arr1, 1)
	fmt.Println(cap(arr1))
	fmt.Println(arr1[1])
	fmt.Println(arr[1])
}
