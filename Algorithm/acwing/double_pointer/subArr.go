package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var xLength, yLength int
	fmt.Fscan(reader, &xLength, &yLength)
	var xArr = make([]int, xLength)
	var yArr = make([]int, yLength)
	for i := 0; i < xLength; i++ {
		fmt.Fscan(reader, &xArr[i])
	}
	for i := 0; i < yLength; i++ {
		fmt.Fscan(reader, &yArr[i])
	}
	var j int
	for i := 0; i < yLength; i++ {
		if j == xLength {
			fmt.Println("Yes")
			return
		}

		if yArr[i] == xArr[j] {
			j++
		}
	}
	if j == xLength {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")

}
