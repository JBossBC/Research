//package main
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//	"strconv"
//)
//
//var graph [][]int
//var n, m, q int
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Fscanln(reader, &n, &m, &q)
//	graph = make([][]int, n)
//	for i := 0; i < n; i++ {
//		graph[i] = make([]int, n)
//	}
//	for i := 0; i < n; i++ {
//		for j := 0; j < n; j++ {
//			if i == j {
//				graph[i][j] = 0
//			} else {
//				graph[i][j] = math.MaxInt
//			}
//		}
//	}
//	for i := 0; i < m; i++ {
//		var x, y, distance int
//		fmt.Fscanln(reader, &x, &y, &distance)
//		x--
//		y--
//		graph[x][y] = int(math.Min(float64(graph[x][y]), float64(distance)))
//	}
//	floyd()
//	var result = make([]string, 0, q)
//	for i := 0; i < q; i++ {
//		var x, y int
//		fmt.Fscanln(reader, &x, &y)
//		if graph[x-1][y-1] > math.MaxInt/2 {
//			result = append(result, "impossible")
//		} else {
//			result = append(result, strconv.Itoa(graph[x-1][y-1]))
//		}
//	}
//	for i := 0; i < len(result); i++ {
//		fmt.Println(result[i])
//	}
//}
//func floyd() {
//	for k := 0; k < n; k++ {
//		for i := 0; i < n; i++ {
//			for j := 0; j < n; j++ {
//				graph[i][j] = int(math.Min(float64(graph[i][j]), float64(graph[i][k]+graph[k][j])))
//			}
//		}
//	}
//}
