package main

import "strings"

func main() {

}
func reverseLeftWords(s string, n int) string {
	var sb = strings.Builder{}
	sb.WriteString(s[n:len(s)])
	sb.WriteString(s[0:n])
	return sb.String()
}
