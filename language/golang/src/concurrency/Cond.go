package main

import "sync"

var lock sync.Mutex

func main() {
	cond := sync.NewCond(&lock)
	for i := 0; i < 100; i++ {
		go func() {
			lock.Lock()
			cond.Broadcast()
		}()
	}
}
