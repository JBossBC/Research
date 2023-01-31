package main

import (
	"fmt"
	"sync"
	"time"
)

var x, y int

const tryTimes = 10000000000

var arr = make([]int64, 14)

func main() {

	now := time.Now()
	group := sync.WaitGroup{}
	group.Add(2)
	go func() {
		defer group.Done()
		for i := 0; i < tryTimes; i++ {
			arr[2] = 1
		}
	}()
	go func() {
		defer group.Done()
		for i := 0; i < tryTimes; i++ {
			arr[1] = 1
		}
	}()
	group.Wait()
	fmt.Println(time.Since(now))
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])

	now2 := time.Now()
	group2 := sync.WaitGroup{}
	group2.Add(2)
	go func() {
		defer group2.Done()
		for i := 0; i < tryTimes; i++ {
			arr[4] = 1
		}
	}()
	go func() {
		defer group2.Done()
		for i := 0; i < tryTimes; i++ {
			arr[8] = 1
		}
	}()
	group2.Wait()
	fmt.Println(time.Since(now2))
	fmt.Println(&arr[4])
	fmt.Println(&arr[8])
}
