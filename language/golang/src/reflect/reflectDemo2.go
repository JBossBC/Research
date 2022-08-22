package main

import (
	"fmt"
	"reflect"
)

type StudentInterface interface {
	getName() string
	getId() int
}
type Student struct {
	Name string
	Id   int
}

func (s *Student) getName() string {
	return s.Name
}
func (s *Student) getId() int {
	return s.Id
}
func main() {
	var number float64 = 1.2222
	var pointer = reflect.ValueOf(&number)
	fmt.Println(pointer)
	var value = reflect.ValueOf(number)
	fmt.Println(value)
	float64Pointer := pointer.Interface().(*float64)
	fmt.Println(*float64Pointer)
	var TypePointer = pointer.Type()
	fmt.Println(TypePointer.Align())
	number = 1.3333
	fmt.Println(reflect.ValueOf(&number))
	var arr = make([]int, 2)
	fmt.Println(reflect.ValueOf(&arr[1]))
	fmt.Println(reflect.ValueOf(number).Float())
	var student = Student{
		Name: "豪猪",
		Id:   1,
	}
	//TODO  i guess the place has implicit conversion,but i hasn't i  enough evidence
	var studentInterfaceObject = reflect.ValueOf(student).Interface()
	fmt.Println(studentInterfaceObject)
	switch studentInterfaceObject.(type) {
	case Student:
		fmt.Println("this is Student struct")
	default:
		fmt.Println("i'm sorry that i can't find this struct ")
	}
	fmt.Println(studentInterfaceObject)
}
