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
	}
	var pointer int
	for index, _ := range discretizationMap {
		convertArr[pointer] = index
		pointer++
	}
	sort.Ints(convertArr)
	for times > 0 {
		var searchBegin, searchEnd = -1, -1
		fmt.Fscan(reader, &searchBegin, &searchEnd)
		var beginPointer, endPointer int = -1, -1
		for j := 0; j < pointer; j++ {
			if convertArr[j] >= searchBegin {
				beginPointer = j
				break
			}
		}
		if beginPointer < 0 {
			if convertArr[0] <= searchEnd {
				beginPointer = 0
			} else {
				continue
			}
		}
		for j := beginPointer; j < pointer; j++ {
			if convertArr[j] > searchEnd {
				endPointer = j - 1
				break
			}
			if j == pointer-1 {
				endPointer = pointer - 1
			}
		}

		var result = 0
		if beginPointer < 0 && endPointer < 0 {
			continue
		}
		for j := beginPointer; j <= endPointer && j >= 0; j++ {
			result += discretizationMap[convertArr[j]]
		}
		fmt.Println(result)
		times--
	}
}
