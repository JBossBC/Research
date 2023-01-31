package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var lock sync.Mutex
//var hello struct {
//	name string
//	lock sync.Mutex
//}

type monitor struct {
	name string
	sex  string
}

// single
var globalMonitor *monitor

var times int64
var locks sync.Mutex
var once sync.Once

//double check lock
func GetMonitor() *monitor {
	////执行到这里的时候
	if globalMonitor == nil {
		locks.Lock()
		//被执行了两次
		//这一步A、B
		if globalMonitor == nil {
			atomic.AddInt64(&times, 1)
			globalMonitor = &monitor{name: "hello"}
		}
		locks.Unlock()
	}
	once.Do(func() {
		globalMonitor = &monitor{name: "xiyang"}
	})
	return globalMonitor
}

var sum111 int64

func main() {
	const times = 100000
	group := sync.WaitGroup{}
	group.Add(times)
	for i := 0; i < times; i++ {
		go func() {
			defer group.Done()
			sum111++
		}()
	}
	group.Wait()
	fmt.Println(sum111)
}

//func main() {
//	const tryTimes = 1000000
//	group := sync.WaitGroup{}
//	group.Add(tryTimes)
//	for i := 0; i < tryTimes; i++ {
//		go func() {
//			defer group.Done()
//			GetMonitor()
//		}()
//	}
//	group.Wait()
//	fmt.Println(times)
//}

//func main() {
//
//	now := time.Now()
//	const retryTimes = 100000
//	group := sync.WaitGroup{}
//	group.Add(retryTimes)
//	for i := 0; i < retryTimes; i++ {
//		go func() {
//			defer group.Done()
//			//time.Sleep(1 * time.Second)
//			ChangeDCL()
//		}()
//	}
//	group.Wait()
//	fmt.Println(time.Since(now))
//	fmt.Println(ExecuteNumber)
//	//fmt.Println(ExecuteNumber)
//}

type Teacher struct {
	Name string
	Id   int
}

var ExecuteNumber int64
var GlobalTeacher *Teacher
var Mutlock sync.Mutex

func DCL() *Teacher {
	Mutlock.Lock()
	if GlobalTeacher == nil {
		atomic.AddInt64(&ExecuteNumber, 1)
		GlobalTeacher = &Teacher{Name: "xiyang"}
	}
	Mutlock.Unlock()
	return GlobalTeacher
}
func ChangeDCL() *Teacher {
	if GlobalTeacher == nil {
		Mutlock.Lock()
		if GlobalTeacher == nil {
			atomic.AddInt64(&ExecuteNumber, 1)
			GlobalTeacher = &Teacher{Name: "xiyang"}
		}
		Mutlock.Unlock()
	}
	return GlobalTeacher
}
