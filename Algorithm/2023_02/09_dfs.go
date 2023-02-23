//package main
//
//import "fmt"
//
//var length int
//var currentPath []int
//var isWalk []bool
//
//func main() {
//	fmt.Scan(&length)
//	currentPath = make([]int, length+1)
//	isWalk = make([]bool, length+1)
//	dfs(0)
//}
//
//func dfs(curLength int) {
//	if length == curLength {
//		for i := 0; i < length; i++ {
//			fmt.Printf("%d", currentPath[i])
//			if i != length-1 {
//				fmt.Printf(" ")
//			}
//		}
//		fmt.Println()
//		return
//	}
//
//	for i := 0; i < length; i++ {
//		if !isWalk[i] {
//			currentPath[curLength] = i + 1
//			isWalk[i] = true
//			dfs(curLength + 1)
//			isWalk[i] = false
//		}
//	}
//
//}
