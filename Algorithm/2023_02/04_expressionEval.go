package main

import (
	"fmt"
	"unicode"
)

var expression string
var nums []int64
var numsPointer int = -1
var operation []byte
var operPointer int = -1

func eval() {
	var first = nums[numsPointer]
	numsPointer--
	var second = nums[numsPointer]
	numsPointer--
	var oper = operation[operPointer]
	operPointer--
	var result int64
	if oper == '+' {
		result = first + second
	} else if oper == '-' {
		result = first - second
	} else if oper == '*' {
		result = first * second
	} else {
		result = first / second
	}
	numsPointer++
	nums[numsPointer] = result
}
func main() {
	fmt.Scanln(&expression)
	operation = make([]byte, len(expression))
	nums = make([]int64, len(expression))
	var mapping = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	for i := 0; i < len(expression); i++ {
		var bytes = expression[i]
		if unicode.IsNumber(rune(bytes)) {
			var number int64 = 0
			var j = i
			for j < len(expression) && unicode.IsNumber(rune(expression[j])) {
				number = number*10 + int64(expression[j]-'0')
				j++
			}
			i = j - 1
			numsPointer++
			nums[numsPointer] = number
		} else if bytes == '(' {
			operPointer++
			operation[operPointer] = bytes
		} else if bytes == ')' {
			for operation[operPointer] != '(' {

				eval()
			}
			operPointer--
		} else {
			for operPointer >= 0 && operation[operPointer] != '(' && mapping[operation[operPointer]] >= mapping[bytes] {
				eval()

			}
			operPointer++
			operation[operPointer] = bytes
		}
	}
	for operPointer >= 0 {
		eval()
	}
	fmt.Println(nums[0])
}
