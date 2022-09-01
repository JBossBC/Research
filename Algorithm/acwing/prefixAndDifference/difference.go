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
	var arr = make([]int, length+1)
	var difference = make([]int, length+2)
	for i := 1; i < length+1; i++ {
		fmt.Fscan(reader, &arr[i])
	}
	for i := 1; i < length; i++ {
		difference[i] = arr[i] - arr[i-1]
	}
	for times > 0 {
		var x, y, value int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		fmt.Fscan(reader, &value)
		difference[x] += value
		difference[y+1] -= value
		times--
	}
	var temp = 0
	for i := 1; i < len(arr); i++ {
		temp += difference[i]
		arr[i] = temp
		fmt.Printf("%d ", arr[i])
	}

}
