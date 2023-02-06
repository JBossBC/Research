package main

func main() {
	findNumberIn2DArray([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5)
}
func findNumberIn2DArray(matrix [][]int, target int) bool {
	var x, y int
	for x < len(matrix) && y < len(matrix[x]) {
		if x+1 < len(matrix) && matrix[x+1][y] < target {
			x = x + 1
		}
		if y+1 < len(matrix[x]) && matrix[x][y+1] < target {
			y = y + 1
		}
		if matrix[x][y] >= target {
			for i := 0; i <= x; i++ {
				if matrix[i][y] == target {
					return true
				}
			}
			for i := 0; i <= y; i++ {
				if matrix[x][i] == target {
					return true
				}
			}
			return false
		}

	}

	return false

}
