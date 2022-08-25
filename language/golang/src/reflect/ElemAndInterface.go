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
	// var pointer interface{} = nil
	// var pointer1 interface{} = nil
	// print(reflect.ValueOf(&pointer))
	// print(reflect.ValueOf(&pointer1))
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
