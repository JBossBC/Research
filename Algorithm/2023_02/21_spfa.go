package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var queue []int
var head, tail int
var n, m int
var graph [][]int
var dist []int
var isExists map[int]bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	queue = make([]int, n*n)
	graph = make([][]int, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		var x, y, distance int
		fmt.Fscanln(reader, &x, &y, &distance)
		x--
		y--
		if graph[x][y] != 0 {
			graph[x][y] = int(math.Min(float64(graph[x][y]), float64(distance)))
			continue
		}
		graph[x][y] = distance
	}
	dist = make([]int, n)
	isExists = make(map[int]bool)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[0] = 0
	isExists[0] = true
	head, tail = 0, -1
	tail++
	queue[tail] = 0
	spfa()
	if dist[n-1] > math.MaxInt/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n-1])
	}
}
func spfa() {
	for tail >= head {
		var point = queue[head]
		for i := 0; i < n; i++ {
			if graph[point][i] == 0 {
				continue
			}
			var distance = graph[point][i] + dist[point]
			if distance < dist[i] {
				dist[i] = distance
				if _, ok := isExists[i]; !ok {
					isExists[i] = true
					tail++
					queue[tail] = i
				}
			}
		}
		head = (head + 1) % len(queue)
		delete(isExists, point)
	}
}
