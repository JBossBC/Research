package main

import "math"

func main() {

}
func minArray(numbers []int) int {
	var minValue int = math.MaxInt
	for i := 0; i < len(numbers); i++ {
		if minValue > numbers[i] {
			minValue = numbers[i]
		}
	}
	return minValue
}
