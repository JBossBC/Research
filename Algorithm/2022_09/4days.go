package main

import "fmt"

func main() {
	var result = numSpecial([][]int{{0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0, 0, 1}, {0, 0, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 1, 0, 0, 0}, {0, 0, 1, 1, 0, 0, 0, 0}})
	fmt.Println(result)
}

//func numSpecial(mat [][]int) int {
//	var line = make([]int, len(mat))
//	var column = make([]int, len(mat[0]))
//	var result int
//	for i := 0; i < len(mat); i++ {
//		for j := 0; j < len(mat[i]); j++ {
//			if mat[i][j] == 1 {
//				if line[i] == 0 && column[j] == 0 {
//					result++
//					line[i] = 1
//					column[j] = 1
//				} else {
//					if line[i] == 1 && column[j] == 1 {
//						result -= 2
//						line[i] = -1
//						column[j] = -1
//					} else if line[i] == 1 {
//						line[i] = -1
//						result--
//					} else if column[j] == 1 {
//						column[j] = -1
//						result--
//					}
//				}
//			}
//		}
//	}
//	return result
//}

func numSpecial(mat [][]int) int {
	var result int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				var tempi = i
				var tempj = j
				for z := 0; z < len(mat[tempi]); z++ {
					if mat[tempi][z] == 1 {
						goto con
					}
				}
				for z := 0; z < len(mat); z++ {
					if mat[z][tempj] == 1 {
						goto con
					}
				}
				result++
			}
		con:
		}
	}
	return result
}
