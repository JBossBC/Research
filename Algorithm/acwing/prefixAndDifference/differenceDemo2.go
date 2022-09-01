package main

import (
	"bufio"
	"fmt"
	"os"
)

var difference [][]int
var arr [][]int

func insert(x int, y int, x1 int, y1 int, value int) {
	difference[x][y] += value
	difference[x1+1][y1+1] += value
	difference[x1+1][y] -= value
	difference[x][y1+1] -= value
}

//TODO
//var tempArr [][]int
//switch i := reflect.ValueOf(a).String(); i {
//case "difference":
//fmt.Println("difference slice is")
//tempArr = difference
//case "arr":
//fmt.Println(" arr slice is")
//tempArr = arr
//default:
//fmt.Println(i)
//panic("type convert error")
//}
//var temp int
//for i := 1; i < len(tempArr); i++ {
//for j := 1; j < len(tempArr[i]); j++ {
//temp = temp + difference[i][j]
//tempArr[i][j] = temp
//fmt.Printf("%d ", tempArr[i][j])
//}
//fmt.Println()
//}
func scanVariable(a interface{}) {
	tempArr := a.([][]int)
	for i := 0; i < len(tempArr); i++ {
		for j := 0; j < len(tempArr[i]); j++ {
			fmt.Printf("%d ", tempArr[i][j])
		}
		fmt.Println()
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	var x, y, times int
	fmt.Fscan(reader, &x, &y, &times)
	arr = make([][]int, x+1)
	difference = make([][]int, x+2)
	difference[len(difference)-1] = make([]int, y+2)
	for i := 0; i < x+1; i++ {
		arr[i] = make([]int, y+1)
		difference[i] = make([]int, y+2)
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			fmt.Fscan(reader, &arr[i][j])
		}
	}
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			insert(i, j, i, j, arr[i][j])
		}
	}
	for times > 0 {
		var a, b, a1, b1, value int
		fmt.Fscan(reader, &a, &b, &a1, &b1, &value)
		insert(a, b, a1, b1, value)
		times--
	}

	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			difference[i][j] += difference[i-1][j] + difference[i][j-1] - difference[i-1][j-1]
			fmt.Printf("%d ", difference[i][j])
		}
		fmt.Println()
	}
}
