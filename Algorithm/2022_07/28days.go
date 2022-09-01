package main

import (
	"fmt"
	"sort"
)

func main() {

	transform := arrayRankTransform([]int{3, 3, 3})
	for i := 0; i < len(transform); i++ {
		fmt.Println(transform[i])
	}
}
func arrayRankTransform(arr []int) []int {
	var tempArr = make([]int, len(arr))
	copy(tempArr, arr)
	sort.Ints(tempArr)
	var resultMap = make(map[int]int, len(arr))
	var mapValue = 0
	for _, value := range tempArr {
		_, ok := resultMap[value]
		if ok {
			continue
		}
		resultMap[value] = mapValue + 1
		mapValue++
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = resultMap[arr[i]]
	}
	return arr
}
