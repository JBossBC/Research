package main

import "strings"

func main() {

}

func fractionAddition(expression string) string {
	figureArr := strings.FieldsFunc(expression, func(r rune) bool {
		return r == '+' || r == '-'
	})
	//单独处理开头
	if figureArr[0] == "-" {
		figureArr = figureArr[1:]
	}

}
