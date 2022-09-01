package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	time := exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"})
	fmt.Println(time)
}

func exclusiveTime(n int, logs []string) []int {
	if n < 1 {
		return nil
	}

	var result = make([]int, n)
	var handlingProcess = make([][]int, n)
	for i := 0; i < len(handlingProcess); i++ {
		handlingProcess[i] = make([]int, 3)
	}
	var arrIndex = -1
	var pointer = 0
	var prePointer = -1
	for pointer < len(logs) {
		split := strings.Split(logs[pointer], ":")
		var index, _ = strconv.Atoi(split[0])
		var symbol = split[1]
		var step, _ = strconv.Atoi(split[2])
		if strings.Compare(symbol, "start") == 0 {
			if arrIndex != -1 {
				prePointer = handlingProcess[arrIndex][0]
				result[prePointer] += step - handlingProcess[arrIndex][1]
			}
			arrIndex++
			handlingProcess[arrIndex][0] = index
			handlingProcess[arrIndex][1] = step
		}
		if strings.Compare(symbol, "end") == 0 {
			result[index] += step - handlingProcess[arrIndex][1]
			if len(handlingProcess) > 1 {
				handlingProcess[prePointer][2] = step
				result[prePointer] -= step - handlingProcess[arrIndex][1]
			}
			arrIndex--
		}
		pointer++
	}
	return result
}
