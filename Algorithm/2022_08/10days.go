package main

import "strings"

func main() {

}
func solveEquation(equation string) string {
	split := strings.Split(equation, "=")
	if len(split[1]) <= 1 {
		return strings.Join(split, "=")
	}
	for i := 0; i < len(split); i++ {

	}
	return ""
}
