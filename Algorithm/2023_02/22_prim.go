package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var graph [][]int
var dist []int
var st []bool
var n, m int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	dist = make([]int, n)
	st = make([]bool, n)
	graph = make([][]int, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		var x, y, distance int
		fmt.Fscanln(reader, &x, &y, &distance)
		if graph[x-1][y-1] != 0 {
			graph[y-1][x-1] = int(math.Min(float64(graph[y-1][x-1]), float64(distance)))
			graph[x-1][y-1] = int(math.Min(float64(graph[x-1][y-1]), float64(distance)))
			continue
		}
		graph[x-1][y-1] = distance
		graph[y-1][x-1] = distance
	}
	result := prim()
}
func prim() int {
	var result int
	for i := 0; i < n; i++ {
		var point int = -1
		for j := 1; j < n; j++ {
			if !st[j] && (point == -1 || dist[point] > dist[j]) {
				point = j
			}
		}
		if i > 0 && dist[i] == 1<<32 {
			return 1 << 32
		}
		if i > 0 {
			result += dist[point]
		}
		st[point] = true
		for j := 1; j < n; j++ {
			dist[j] = int(math.Min(float64(dist[j]), float64(graph[point][j])))
		}
	}
	return result
}
