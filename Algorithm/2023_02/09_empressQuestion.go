package main

import "fmt"

var chessboard [][]int
var n int
var col, dg, udg []bool

func main() {
	fmt.Scan(&n)
	chessboard = make([][]int, n)
	col = make([]bool, n*2)
	dg = make([]bool, n*2)
	udg = make([]bool, n*2)
	for i := 0; i < len(chessboard); i++ {
		chessboard[i] = make([]int, n)
	}
	empressDFS(0)
}

func printEmpress() {
	for i := 0; i < len(chessboard); i++ {
		for j := 0; j < len(chessboard[i]); j++ {
			if chessboard[i][j] == 1 {
				fmt.Print("Q")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func empressDFS(curNumber int) {
	if curNumber == n {
		printEmpress()
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		y := curNumber
		if dg[i-y+n] == false && col[i] == false && udg[y+i] == false {
			col[i] = true
			dg[i-y+n] = true
			udg[y+i] = true
			chessboard[y][i] = 1
			empressDFS(curNumber + 1)
			chessboard[y][i] = 0
			col[i] = false
			dg[i-y+n] = false
			udg[y+i] = false
		}
	}
}
