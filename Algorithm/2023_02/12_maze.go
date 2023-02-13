//package main
//
//import "fmt"
//
//type pair struct {
//	first, second int
//}
//
//var g [][]int
//var d [][]int
//var n, m int
//
//func main() {
//	fmt.Scanf("%d%d", &n, &m)
//	g = make([][]int, n)
//	d = make([][]int, n)
//	for i := 0; i < n; i++ {
//		g[i] = make([]int, m)
//		d[i] = make([]int, m)
//		for j := 0; j < m; j++ {
//			fmt.Scanf("%d", &g[i][j]) // 读入地图
//			d[i][j] = -1              // 初始化为 -1 表示没有走过
//		}
//	}
//	d[0][0] = 0 // 表示0，0位置已经走过了
//
//	fmt.Printf("%d", bfs())
//
//}
//func bfs() int {
//	q := make([]pair, 0)
//	q = append(q, pair{0, 0})
//	dx := []int{-1, 0, 1, 0}
//	dy := []int{0, 1, 0, -1}
//	for len(q) > 0 {
//		t := q[0]
//		q = q[1:]
//		a, b := t.first, t.second
//		for i := 0; i < 4; i++ {
//			lx := a + dx[i]
//			ly := b + dy[i]
//			if lx >= 0 && lx < n && ly >= 0 && ly < m && g[lx][ly] == 0 && d[lx][ly] == -1 {
//				d[lx][ly] = d[a][b] + 1
//				q = append(q, pair{lx, ly})
//			}
//		}
//
//	}
//	return d[n-1][m-1]
//}
