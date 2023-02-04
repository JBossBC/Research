package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 6
	var y int = 8
	fmt.Println(&x)
	fmt.Println(y)
	alignValue := unsafe.Add(unsafe.Pointer(&x), 3*unsafe.Alignof(x))
	alignValue = unsafe.Pointer(&y)
	fmt.Println(*alignValue)
	reflect.

}
