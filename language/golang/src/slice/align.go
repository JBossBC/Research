package main

import "fmt"

func main() {
	//var x int = 6
	//var y int = 8
	//fmt.Println(&x)
	//fmt.Println(y)
	//alignValue := unsafe.Add(unsafe.Pointer(&x), 3*unsafe.Alignof(x))
	//alignValue = unsafe.Pointer(&y)
	//fmt.Println(*alignValue)
	var slice = []int{1, 1, 2, 3, 4, 5}
	var temp = slice[0:0]
	temp = append(temp, 2)
	fmt.Println(temp)
}
