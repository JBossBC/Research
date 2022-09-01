package main

import "math"

//TODO
func main() {

}

func maximalRectangle(matrix [][]byte) int {
	var SelectSquare = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {

		}
	}
}

func SquareArea(i int, j int, arr [][]byte) int {
	var routeDynimic = make([][]byte, 2)
	for i := 0; i < len(routeDynimic); i++ {
		routeDynimic[i] = make([]byte, int(math.Max(float64(len(arr)), float64(len(arr[0])))))
	}

}
