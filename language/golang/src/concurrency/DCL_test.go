package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go DCL()
	}
}
func BenchmarkChangeDCL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go ChangeDCL()
	}
}

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
		atomic.AddInt64(&ExecuteNumber, 1)
		GlobalTeacher = &Teacher{Name: "xiyang"}
		Mutlock.Unlock()
	}
	return GlobalTeacher
}
