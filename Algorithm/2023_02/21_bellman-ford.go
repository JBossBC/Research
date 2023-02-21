//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//)
//
//type edge struct {
//	a, b, c int
//}
//
//var dist []int
//var last []int
//
////需要遍历所有边
//var graph []edge
//var n, m, k int
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Fscanln(reader, &n, &m, &k)
//	graph = make([]edge, 0, m)
//	dist = make([]int, n)
//	last = make([]int, n)
//	for i := 0; i < m; i++ {
//		var a, b, c int
//		fmt.Fscanln(reader, &a, &b, &c)
//		graph = append(graph, edge{a: a - 1, b: b - 1, c: c})
//	}
//	bellman_ford()
//	if dist[n-1] > math.MaxInt/2 {
//		fmt.Println("impossible")
//	} else {
//		fmt.Println(dist[n-1])
//	}
//}
//func bellman_ford() {
//	for i := 0; i < len(dist); i++ {
//		dist[i] = math.MaxInt
//	}
//	dist[0] = 0
//	for i := 0; i < k; i++ {
//		copy(last, dist)
//		for j := 0; j < m; j++ {
//			var tempEdge = graph[j]
//			dist[tempEdge.b] = int(math.Min(float64(dist[tempEdge.b]), float64(last[tempEdge.a]+tempEdge.c)))
//		}
//	}
}
