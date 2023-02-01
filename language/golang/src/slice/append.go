package main

import "fmt"

func main() {
	slice := make([]int, 0)
	for i := 0; i < 100000; i++ {
		slice = append(slice, 10)
		slice[i] = 0
	}
	math.M
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d", slice[i])

	}
	fmt.Println("-----------------------")
	fmt.Println(len(slice))
	fmt.Println("-------------")
	fmt.Println(cap(slice))
}
