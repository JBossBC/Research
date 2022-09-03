package main

import "sort"

func main() {
	findLongestChain([][]int{{-6, 9}, {1, 6}, {8, 10}, {-1, 4}, {-6, -2}, {-9, 8}, {-5, 3}, {0, 3}})
}
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a][0] < pairs[b][0]
	})
	var result = 1
	var temp = pairs[0]
	for i := 1; i < len(pairs); i++ {
		if pairs[i][1] <= temp[1] {
			temp[1] = pairs[i][1]
		}
		if pairs[i][0] > temp[1] {
			result++
			temp = pairs[i]
		}
	}
	return result
}
