package main

import (
	"fmt"
	"reflect"
)

func main() {
	var student = Student{
		Id:   1,
		name: "xiyang",
	}
	createQuery(student)
}

type Student struct {
	Id   int
	name string
}

//对于任意结构体创建Sql语句
func createQuery(q interface{}) string {
	//判断类型为结构体
	var query string
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query = fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s,\"%s\"", query, v.Field(i).String())
				}
			default:
				panic("the type can't support reflect,please connect with the coder adding this reflect supporting")
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
	}
	// if !query.(string) {
	// 	return nil
	// }
	return query
}
