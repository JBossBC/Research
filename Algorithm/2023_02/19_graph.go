package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var r, s []int
var sIndex int
var times, pn int
var graph [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &times, &pn)
	graph = make([][]int, pn)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, pn)
	}
	sIndex = -1
	r = make([]int, pn)
	s = make([]int, pn)
	for i := 0; i < pn; i++ {
		r[i] = math.MaxInt
	}
	r[0] = 0
	for i := 0; i < times; i++ {
		var x, y, length int
		fmt.Fscanln(reader, &x, &y, &length)
		if x == y {
			continue
		}
		x--
		y--
		if graph[x][y] != 0 {
			graph[x][y] = int(math.Min(float64(graph[x][y]), float64(length)))
			continue
		}
		graph[x][y] = length
	}
	sIndex++
	s[sIndex] = 0
	dijkstra()
	for i := 0; i < pn; i++ {
		println(r[i])
	}
}
func dijkstra() {
	for sIndex >= 0 {
		var temp = s[sIndex]
		sIndex--
		for i := 1; i < pn; i++ {
			if graph[temp][i] == 0 || temp == i {
				continue
			}
			r[i] = int(math.Min(float64(r[temp]+graph[temp][i]), float64(r[i])))
			sIndex++
			s[sIndex] = i
		}

	}
	sort.Ints(r)

}
