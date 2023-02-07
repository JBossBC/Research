package main

import (
	"bufio"
	"fmt"
	"os"
)

var trieArr [][26]int

var hasStr map[int]int
var pointer = 0

func insert(str []byte) {
	var result int
	for i := 0; i < len(str); i++ {
		var numsConv = str[i] - 'a'
		if trieArr[i][numsConv] == 0 {
			pointer++
			trieArr[i][numsConv] = pointer
		}
	}
	result = trieArr[len(str)-1][str[len(str)-1]-'a']
	hasStr[result]++
}
func query(str []byte) {
	var result int
	for i := 0; i < len(str); i++ {
		var numsConv = str[i] - 'a'
		if trieArr[i][numsConv] == 0 {
			resultOutput = append(resultOutput, 0)
			return
		}
		result = trieArr[i][numsConv]
	}
	result = trieArr[len(str)-1][str[len(str)-1]-'a']
	if value, ok := hasStr[result]; ok {
		resultOutput = append(resultOutput, value)
	} else {
		resultOutput = append(resultOutput, 0)
	}

}

var resultOutput = make([]int, 0, 10010)

func main() {
	hasStr = make(map[int]int)
	trieArr = make([][26]int, 100100)
	reader := bufio.NewReader(os.Stdin)
	var operationNums int
	fmt.Fscanln(reader, &operationNums)
	for i := 0; i < operationNums; i++ {
		var operation byte
		var str string
		fmt.Fscanf(reader, "%c %s\n", &operation, &str)
		switch operation {
		case 'I':
			insert([]byte(str))
		case 'Q':
			query([]byte(str))
		}
	}
	for i := 0; i < len(resultOutput); i++ {
		fmt.Println(resultOutput[i])
	}
}
