package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var length, times int
	fmt.Fscan(reader, &length, &times)
	var temp = 0
	var arr = make([]int, length)
	var prefix = make([]int, length+1)
	for i := 0; i < length; i++ {
		fmt.Fscan(reader, &arr[i])
		prefix[i+1] = temp + arr[i]
		temp += arr[i]
	}
	for times > 0 {
		var l, r int
		fmt.Fscan(reader, &l, &r)
		fmt.Println(prefix[r] - prefix[l-1])
		times--
	}
}
