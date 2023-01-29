package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x, y int
	lock sync.Mutex
)

func main() {
	count := 0
	pre := time.Now()
	for count < 100000 {
		a, b := 0, 0
		count++
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			lock.Lock()
			a = 1
			x = b
			lock.Unlock()
			wg.Done()
		}()
		go func() {
			lock.Lock()
			b = 1
			y = a
			lock.Unlock()
			wg.Done()
		}()
		wg.Wait()
		if x == 0 && y == 0 {
			fmt.Println("error", x, y)
			return
		}
	}
	fmt.Println("you success")
	println(time.Since(pre).String())
}
