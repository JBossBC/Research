package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var arr = make([]int, 10)
	var arrOrd = [4]int{1, 2, 3}
	sizeof := unsafe.Sizeof(&arr)
	fmt.Println(unsafe.Pointer(&arr))
	add := unsafe.Add(unsafe.Pointer(&arr), 12)
	fmt.Println(unsafe.Alignof(&arrOrd))
	fmt.Println(sizeof)
	fmt.Println(unsafe.Pointer(&arr))
	fmt.Println(add)
	fmt.Println(cap(arr))
	fmt.Println(unsafe.Sizeof(add))
	fmt.Println(unsafe.Pointer(&arr))
	arrPointerSize := &arr
	fmt.Println(unsafe.Pointer(arrPointerSize))
}
