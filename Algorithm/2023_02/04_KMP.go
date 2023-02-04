package main

import (
	"bufio"
	"fmt"
	"os"
)

var next = make([]int, 100100)

func main() {
	var subLength, originLength int
	var subStr, originStr string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &subLength)
	fmt.Fscanln(reader, &subStr)
	fmt.Fscanln(reader, &originLength)
	fmt.Fscanln(reader, &originStr)
	//build index
	var begin int = 0
	for i := 2; i < subLength; i++ {
		for begin != 0 && subStr[i] != subStr[begin+1] {
			begin = next[begin]
		}
		if subStr[i] == subStr[begin+1] {
			begin++
		}
		next[i] = begin
	}
	begin = 0
	for i := 1; i < originLength; i++ {
		for begin != 0 && originStr[i] != subStr[begin+1] {
			begin = next[begin]
		}
		if originStr[i] == subStr[begin+1] {
			begin++
		}
		if begin == subLength {
			fmt.Printf("%d ", i-subLength+1)
			begin = next[begin]
		}
	}

}
