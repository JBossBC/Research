package main

import "sort"

func main() {

}

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)

	r := 1
	for _, coin := range coins {
		if coin > r {
			break
		}
		r += coin
	}

	return r
}
