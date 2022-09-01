package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//超时
func main() {
	reader := bufio.NewReader(os.Stdin)
	var length, times int
	fmt.Fscan(reader, &length, &times)
	var discretizationMap = make(map[int]int, length)
	var index, value int
	var convertArr = make([]int, length)
	//离散化
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &index, &value)
		discretizationMap[index] += value
		convertArr[i] = index
	}
	sort.Ints(convertArr)
	var unequalArr = make([]int, length)
	var pointerIndex = 0
	//去掉重复的元素,double pointer
	var i = 0
	for i < length {
		var temp = i + 1
		for temp < length && convertArr[temp] == convertArr[i] {
			temp++
		}
		unequalArr[pointerIndex] = convertArr[i]
		i = temp
		pointerIndex++
	}
	pointerIndex--
	for times > 0 {
		var searchBegin, searchEnd = -1, -1
		fmt.Fscan(reader, &searchBegin, &searchEnd)
		var tempArr = make([]int, 0)
		for j := 0; j < len(unequalArr); j++ {
			if unequalArr[j] >= searchBegin && unequalArr[j] <= searchEnd {
				tempArr = append(tempArr, unequalArr[j])
			}
			if unequalArr[j] >= searchEnd {
				break
			}
		}
		var result = 0
		for j := 0; j < len(tempArr); j++ {
			result += discretizationMap[tempArr[j]]
		}
		fmt.Println(result)
		times--
	}
}
