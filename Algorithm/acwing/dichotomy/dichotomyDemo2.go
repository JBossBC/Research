package main

import (
	"fmt"
)

func main() {
	var number float64
	fmt.Scan(&number)
	f := dichotomy(number)
	fmt.Printf("%f", f)
}

func dichotomy(number float64) float64 {
	var l float64 = -10000
	var r float64 = 10000
	var middle = (l + r) / 2
	for r-l >= 1e-7 {
		middle = (l + r) / 2
		if middle*middle*middle < number {
			l = middle
		} else {
			r = middle
		}
	}
	return r
}
