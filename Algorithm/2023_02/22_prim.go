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
	dist = make([]int, n+1)
	st = make([]bool, n+1)
	graph = make([][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		var x, y, distance int
		fmt.Fscanln(reader, &x, &y, &distance)
		if graph[x][y] != 0 {
			graph[y][x] = int(math.Min(float64(graph[y][x]), float64(distance)))
			graph[x][y] = int(math.Min(float64(graph[x][y]), float64(distance)))
			continue
		}
		graph[x][y] = distance
		graph[y][x] = distance
	}
	for i := 0; i < len(dist); i++ {
		dist[i] = 1 << 31
	}
	dist[1] = 1
	result := prim()
	if result == 1<<31 {
		fmt.Println("impossible")
	} else {
		fmt.Println(result)
	}
}
func prim() int {
	var result int
	for i := 0; i < n; i++ {
		var point int = -1
		for j := 1; j <= n; j++ {
			if !st[j] && (point == -1 || dist[point] > dist[j]) {
				point = j
			}
		}
		if i > 0 && dist[point] == 1<<31 {
			return 1 << 31
		}
		if i > 0 {
			result += dist[point]
		}
		st[point] = true
		for j := 1; j <= n; j++ {
			if graph[point][j] != 0 {
				dist[j] = int(math.Min(float64(dist[j]), float64(graph[point][j])))
			}
		}
	}
	return result
}
