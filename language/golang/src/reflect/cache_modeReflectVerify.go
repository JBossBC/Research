package main

import (
	"fmt"
)

const number1 = 11
const number2 = 11

func main() {
	var variableNumber = 11
	var variableNumberEqual = 11
	// var variableNumber1 = 555555
	// var variableNumber1Equal = 555555
	// this is  evidence that the var between -128 and 127 hasn't use the cache mode
	fmt.Println("this is variableNumber pointer values ", &variableNumber)
	fmt.Println("this is variableNumber pointer values ", &variableNumberEqual)

	//TODO const ???
	fmt.Println("const number1 address is", &number1)
	fmt.Println("const number2 address is", &number2)
}
