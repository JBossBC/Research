package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

//var lock sync.Mutex

var sum int64

func init() {
	sum = 1
}

type student struct {
	name string
	id   int
	sex  int64
}

var globalStudent *student
var localInt int64

func main() {
	create, _ := os.Create("trace")
	group := sync.WaitGroup{}
	trace.Start(create)
	group.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer group.Done()
			localInt++
		}()
	}
	group.Wait()
	fmt.Println(localInt)
	runtime.GC()

	trace.Stop()
	//fmt.Println(i)
	//pprof.WriteHeapProfile(create)
}

//
//func getObj() *student {
//	if globalStudent == nil {
//		lock.Lock()
//		globalStudent = &student{
//			name: "hello",
//			id:   0,
//			sex:  localInt,
//		}
//		lock.Unlock()
//	}
//	return globalStudent
//}
