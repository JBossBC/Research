package main

import (
	"fmt"
)

func main() {
	people := groupThePeople([]int{3, 3, 3, 3, 3, 1, 3})
	for _, person := range people {
		fmt.Println(person)
	}
}
func groupThePeople(groupSizes []int) [][]int {
	var tempMap = make(map[int][]int, 0)
	for i := 0; i < len(groupSizes); i++ {
		if tempMap[groupSizes[i]] == nil {
			tempMap[groupSizes[i]] = make([]int, 0, 100)
		}
		tempMap[groupSizes[i]] = append(tempMap[groupSizes[i]], i)
	}
	var result = make([][]int, 0, 100)
	var tempSum = 0
	for index, value := range tempMap {
		var arr = make([]int, 0)
		for i := 0; i < len(value); i++ {
			arr = append(arr, value[i])
			tempSum++
			if tempSum == index {
				result = append(result, arr)
				arr = make([]int, 0)
				tempSum = 0
			}
		}
	}
	return result

}
