package main

import (
	"bufio"
	"fmt"
	"os"
)

//暴力解法!!!
// func main() {
// 	var xLength, yLength, target int
// 	fmt.Scan(&xLength, &yLength, &target)
// 	var xArr = make([]int, xLength)
// 	var yArr = make([]int, yLength)
// 	for i := 0; i < xLength; i++ {
// 		fmt.Scan(&xArr[i])
// 	}
// 	for i := 0; i < yLength; i++ {
// 		fmt.Scan(&yArr[i])
// 	}
// 	var result = make(map[int][]int,xLength*yLength)
// 	var pointer=0
// 	for i := 0; i < xLength; i++ {
// 		var j = 0
// 		if xArr[i]+yArr[j]>target{
// 			break
// 		}
// 		for j < yLength && xArr[i]+yArr[j] <= target {
// 			if xArr[i]+yArr[j] == target {
// 				result[pointer]=[]int{i,j}
// 				pointer++
// 			}
// 			j++
// 		}
// 	}
// 	for i:=0;i<pointer;i++{
// 		fmt.Printf("%d %d \n",result[i][0],result[i][1])
// 	}
// }

//双指针的步骤:先求出暴力解法，再观察其有没有单调性，再进行优化(能否用双指针需要考虑他的单调性)
// func main() {
// 	var xLength, yLength, target int
// 	fmt.Scan(&xLength, &yLength, &target)
// 	var xArr = make([]int, xLength)
// 	var yArr = make([]int, yLength)
// 	for i := 0; i < xLength; i++ {
// 		fmt.Scan(&xArr[i])
// 	}
// 	for i := 0; i < yLength; i++ {
// 		fmt.Scan(&yArr[i])
// 	}
// 	var resultx, resulty = -1, -1
// 	var j = yLength - 1
// 	for i := 0; i < xLength; i++ {
// 		for j >= 0 && xArr[i]+yArr[j] > target {
// 			j--
// 		}
// 		if j >= 0 && xArr[i]+yArr[j] == target {
// 			//resultx = i
// 			//resulty = j
// 			fmt.Printf("%d %d", i, j)
// 			break
// 		}
// 	}
// 	fmt.Printf("%d %d", resultx, resulty)
// }

func main() {
	reader := bufio.NewReader(os.Stdin)
	var xLength, yLength, target int
	fmt.Fscan(reader, &xLength, &yLength, &target)
	var xArr = make([]int, xLength)
	var yArr = make([]int, yLength)
	for i := 0; i < xLength; i++ {
		fmt.Fscan(reader, &xArr[i])
	}
	for i := 0; i < yLength; i++ {
		fmt.Fscan(reader, &yArr[i])
	}
	var j = yLength - 1
	for i := 0; i < xLength; i++ {
		for j >= 0 && xArr[i]+yArr[j] > target {
			j--
		}
		if j >= 0 && xArr[i]+yArr[j] == target {
			fmt.Printf("%d %d", i, j)
			return
		}
	}

}
