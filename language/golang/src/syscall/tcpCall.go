package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var i int32 = 0
	var barrier int32
	go func() {
		for atomic.LoadInt32(&i) == 1 {
			fmt.Println(i)
		}
	}()
	go func() {
		for {
			atomic.StoreInt32(&barrier, 1)
			atomic.StoreInt32(&i, 1)
			atomic.LoadInt32(&barrier)
		}
	}()
	time.Sleep(10 * time.Second)
	fmt.Println(i)

}
