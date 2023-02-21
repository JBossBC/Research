//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//)
//
//var pn, times int
//var graph [][]int
//var s []int
//var heap PairHeap
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Fscanln(reader, &pn, &times)
//	graph = make([][]int, pn)
//	for i := 0; i < len(graph); i++ {
//		graph[i] = make([]int, pn)
//	}
//	s = make([]int, pn)
//	for i := 0; i < times; i++ {
//		var x, y, length int
//		fmt.Fscanln(reader, &x, &y, &length)
//		if graph[x-1][y-1] != 0 {
//			graph[x-1][y-1] = int(math.Min(float64(graph[x-1][y-1]), float64(length)))
//			continue
//		}
//		graph[x-1][y-1] = length
//	}
//	heap = PairHeap{}
//	heap.Push(pair{point: 0, distance: 0})
//	for i := 1; i < pn; i++ {
//		heap.Push(pair{point: i, distance: math.MaxInt})
//	}
//	for i := 0; i < len(s); i++ {
//		s[i] = math.MaxInt
//	}
//	s[0] = 0
//	dijkstra()
//}
//func dijkstra() {
//	for heap.Len() > 0 {
//		var temp = heap.Pop().(pair)
//		tempP, _ := temp.point, temp.distance
//		for i := 0; i < pn; i++ {
//			var newValue = s[tempP] + graph[tempP][i]
//			if graph[tempP][i] == 0 || newValue > s[i] {
//				continue
//			}
//			s[i] = newValue
//		}
//	}
//}
//
//type pair struct {
//	point    int
//	distance int
//}
//type PairHeap []pair
//
//func (p PairHeap) Len() int {
//	return len(p)
//}
//func (p PairHeap) Less(i, j int) bool {
//	return p[i].distance < p[j].distance
//}
//func (p PairHeap) Swap(i, j int) {
//	p[i].point, p[j].point = p[j].point, p[i].point
//	p[i].distance, p[j].distance = p[j].distance, p[i].distance
//}
//func (p *PairHeap) Push(x any) {
//	*p = append(*p, x.(pair))
//}
//func (p *PairHeap) Pop() any {
//	old := *p
//	n := len(old)
//	x := old[n-1]
//	*p = old[0 : n-1]
//	return x
//}
