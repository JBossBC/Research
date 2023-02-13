package main

import (
	"fmt"
	"unsafe"
)

type copyCheck uintptr

type nocopy struct{}
type hello struct {
	copyCheck copyCheck
	nocopy    nocopy
}
type hello2 struct {
	copyCheck uintptr
}

func main() {
	//h := hello{}
	//fmt.Println(uintptr(h.copyCheck))
	//fmt.Println(uintptr(unsafe.Pointer(&h.copyCheck)))
	//atomic.CompareAndSwapUintptr((*uintptr)(&h.copyCheck), 0, uintptr(unsafe.Pointer(&h.copyCheck)))
	//y := h
	//fmt.Println(uintptr(h.copyCheck))
	//atomic.CompareAndSwapUintptr((*uintptr)(&h.copyCheck), 0, uintptr(unsafe.Pointer(&h.copyCheck)))
	//fmt.Println(uintptr(unsafe.Pointer(&y.copyCheck)))
	h := &hello2{}
	//fistUse
	h.copyCheck = uintptr(unsafe.Pointer(&h))

	g := h
	if !(h.copyCheck == uintptr(unsafe.Pointer(&g))) {
		//已经被复制
		fmt.Println("已经被复制")
	}

}
