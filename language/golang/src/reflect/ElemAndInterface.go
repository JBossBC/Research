package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var pointer = &Student{
		Name: "豪猪",
		Id:   1,
	}
	var x interface{} = pointer
	var reflectValue = reflect.ValueOf(x).Elem()
	// var printString = fmt.Sprint("this is Student Name %s and  Id %d by reflect", reflectValue.Name, reflectValue.Id)
	// fmt.Println(printString)
	fmt.Println(reflectValue.Kind())
	fmt.Println(reflectValue.Elem().Kind())
	fmt.Println(reflectValue.Elem().Elem().Kind())
}
