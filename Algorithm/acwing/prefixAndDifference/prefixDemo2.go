package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var Xlength, Ylength, times int
	fmt.Fscan(reader, &Xlength, &Ylength, &times)
	var arr = make([][]int, Xlength)
	var prefix = make([][]int, Xlength+1)
	prefix[0] = make([]int, Ylength+1)
	for i := 0; i < Xlength; i++ {
		arr[i] = make([]int, Ylength)
		prefix[i+1] = make([]int, Ylength+1)
	}
	for i := 0; i < Xlength; i++ {
		for j := 0; j < Ylength; j++ {
			fmt.Fscan(reader, &arr[i][j])
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] + arr[i][j] - prefix[i][j]
		}
	}
	for times > 0 {
		var x, y, x1, y1 int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)
		fmt.Fscan(reader, &x1)
		fmt.Fscan(reader, &y1)
		fmt.Println(prefix[x1][y1] - prefix[x-1][y1] - prefix[x1][y-1] + prefix[x-1][y-1])
		times--
	}
}
