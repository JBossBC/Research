package main

func main() {
	//print(findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20))
	//print(findNumberIn2DArray([][]int{{-5}}, -2))
	//print(findNumberIn2DArray([][]int{{1, 1}}, 0))
	print(findNumberIn2DArray([][]int{{5, 6, 10, 14}, {6, 10, 13, 18}, {10, 13, 18, 19}}, 14))
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false

}
