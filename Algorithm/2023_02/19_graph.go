package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var s map[int]bool
var dist []int
var times, pn int
var graph [][]int

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &pn, &times)
	graph = make([][]int, pn)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, pn)
	}
	for i := 0; i < times; i++ {
		var x, y, distance int
		fmt.Fscanln(reader, &x, &y, &distance)
		if graph[x-1][y-1] != 0 {
			graph[x-1][y-1] = int(math.Min(float64(distance), float64(graph[x-1][y-1])))
			continue
		}
		graph[x-1][y-1] = distance
	}
	s = make(map[int]bool)
	dist = make([]int, pn)
	for i := 0; i < pn; i++ {
		dist[i] = math.MaxInt
	}
	dist[0] = 0
	dijkstra()
	if dist[len(dist)-1] == math.MaxInt {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[len(dist)-1])
	}
}
func dijkstra() {
	for i := 0; i < pn; i++ {
		var minIndex int
		var value int = math.MaxInt
		for j := 0; j < pn; j++ {
			if !s[j] && dist[j] < value {
				minIndex = j
				value = dist[j]
			}
		}
		s[minIndex] = true
		for j := 0; j < pn; j++ {
			if j == minIndex || graph[minIndex][j] == 0 || dist[j] < dist[minIndex]+graph[minIndex][j] {
				continue
			}
			dist[j] = dist[minIndex] + graph[minIndex][j]
		}
	}
}
