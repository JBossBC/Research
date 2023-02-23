package main

import "fmt"

func main() {
	var MaxTimes = 1000
	for j := 0; j < MaxTimes; j++ {
		var result = circularPermutation(3, 5)
		var answer = []int{5, 4, 0, 1, 3, 2, 6, 7}
		var flag = true
		for i := 0; i < len(result); i++ {
			if answer[i] != result[i] {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println("success")
			for i := 0; i < len(result); i++ {
				fmt.Println(result[i])
			}
		}
	}
}

var result []int

// var queue []int
// var head,tail int
var head int
var exist map[int]bool

func circularPermutation(n int, start int) []int {
	result = make([]int, 2<<(n-1))
	head = 0
	// queue =make([]int,n)
	//     head,tail =0,-1
	//    tail++
	//    queue[tail]=start
	exist = make(map[int]bool)
	exist[start] = true
	result[head] = start
	head++
	dfs(start, 2<<(n-1))
	return result

}
func dfs(temp int, n int) bool {
	var matching map[int]bool = make(map[int]bool)
	for i := 0; i < n; i++ {
		if exist[i] || !isMatching(temp, i) {
			continue
		}
		matching[i] = true
	}
	if len(matching) == 0 {
		return false
	}
	for index, _ := range matching {
		exist[index] = true
		result[head] = index
		head++
		if dfs(index, n) && head == len(result) && isMatching(result[0], result[head-1]) {
			return true
		}
		head--
		exist[index] = false
	}

	return false
}
func isMatching(x int, y int) bool {
	var diff int
	for x != 0 || y != 0 {
		var tempX = x % 2
		var tempY = y % 2
		if tempX != tempY {
			if diff != 1 {
				diff++
			} else {
				return false
			}
		}
		x = x / 2
		y = y / 2
	}
	if diff == 1 {
		return true
	}
	return false
}
