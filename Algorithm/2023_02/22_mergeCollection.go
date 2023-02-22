package main

import (
	"bufio"
	"fmt"
	"os"
)

var collection []int
var result []string

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	collection = make([]int, n)
	result = make([]string, 0, m)
	for i := 0; i < len(collection); i++ {
		collection[i] = i
	}
	for i := 0; i < m; i++ {
		var opr byte
		var x, y int
		fmt.Fscanln(reader, &opr, &x, &y)
		switch opr {
		case 'M':
			merge(x, y)
		case 'Q':
			query(x, y)
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
func find(a int) int {
	if a != collection[a] {
		collection[a] = find(collection[a])
	}
	return collection[a]
}
func merge(a int, b int) {
	collection[find(b)] = find(a)
}
func query(a int, b int) {
	if find(a) == find(b) {
		result = append(result, "Yes")
	} else {
		result = append(result, "No")
	}
}
