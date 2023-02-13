//package main
//
//func main() {
//	println(alphabetBoardPath("code"))
//}
//
//var route []byte
//var findIndex int = 0
//var mapping = map[int]byte{0: 'U', 1: 'D', 2: 'L', 3: 'R'}
//var targetIndex int = 0
//var routeLength = -1
//var board = [][]byte{{'a', 'b', 'c', 'd', 'e'}, {'f', 'g', 'h', 'i', 'j'}, {'k', 'l', 'm', 'n', 'o'}, {'p', 'q', 'r', 's', 't'}, {'u', 'v', 'w', 'x', 'y'}, {'z'}}
//var result string
//var find = false
//var hasWalk [][]bool
//
//func alphabetBoardPath(target string) string {
//	route = make([]byte, 1000)
//	targetIndex = len(target)
//	result = target
//	routeLength = -1
//	findIndex = 0
//	hasWalk = make([][]bool, len(board))
//	for i := 0; i < len(hasWalk); i++ {
//		hasWalk[i] = make([]bool, len(board[i]))
//	}
//	find = false
//	bfs(0, 0)
//	return string(route[:routeLength+1])
//}
//
//func bfs(x int, y int) {
//	if board[x][y] == result[findIndex] {
//		findIndex++
//		routeLength++
//		route[routeLength] = '!'
//	}
//	if findIndex == targetIndex {
//		find = true
//		return
//	}
//	hasWalk[x][y] = true
//	for i := 0; i < 4; i++ {
//		var tempRoute = mapping[i]
//		routeLength++
//		route[routeLength] = tempRoute
//		switch tempRoute {
//		case 'U':
//			if x-1 >= 0 && (!hasWalk[x-1][y]) {
//				bfs(x-1, y)
//			}
//		case 'D':
//			if x+1 < len(board) && y < len(board[x+1]) && (!hasWalk[x+1][y]) {
//				bfs(x+1, y)
//			}
//		case 'L':
//			if y-1 >= 0 && (!hasWalk[x][y-1]) {
//				bfs(x, y-1)
//			}
//		case 'R':
//			if x < len(board) && y+1 < len(board[x]) && (!hasWalk[x][y+1]) {
//				bfs(x, y+1)
//			}
//		}
//		routeLength--
//	}
//	hasWalk[x][y] = false
//}
