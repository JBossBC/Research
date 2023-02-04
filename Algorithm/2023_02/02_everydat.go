package main

//func main() {
//	arr := shortestAlternatingPaths(5, [][]int{{2, 2}, {0, 1}, {0, 3}, {0, 0}, {0, 4}, {2, 1}, {2, 0}, {1, 4}, {3, 4}}, [][]int{{1, 3}, {0, 0}, {0, 3}, {4, 2}, {1, 0}})
//	for i := 0; i < len(arr); i++ {
//		fmt.Println(arr[i])
//	}
//}

//func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
//	//
//	wayRoute := make([][]int, n)
//	var result = make([]int, n)
//	for i := 0; i < len(wayRoute); i++ {
//		wayRoute[i] = make([]int, n)
//	}
//	for i := 0; i < len(redEdges); i++ {
//		var x = redEdges[i][0]
//		var y = redEdges[i][1]
//		wayRoute[x][y] = 1
//	}
//	for i := 0; i < len(blueEdges); i++ {
//		var x = blueEdges[i][0]
//		var y = blueEdges[i][1]
//		if wayRoute[x][y] != 1 {
//			wayRoute[x][y] = -1
//		} else {
//			wayRoute[x][y] = 2
//		}
//	}
//	for i := 1; i < n; i++ {
//		if wayRoute[0][i] != 0 {
//			result[i] = 1
//			var consistentArrive = make(map[int]int)
//			consistentArrive[i] = wayRoute[0][i]
//			var wayLength = make(map[int]int)
//			wayLength[i] = 1
//
//			for len(consistentArrive) > 0 {
//				var currentBase int
//				var preSelect int
//				for key, value := range consistentArrive {
//					preSelect = value
//					currentBase = key
//					break
//				}
//				for j := 1; j < n; j++ {
//					if j == currentBase {
//						continue
//					}
//					if wayRoute[currentBase][j] != 0 {
//						if (wayRoute[currentBase][j] ^ preSelect) == -2 {
//						} else if wayRoute[currentBase][j] == 2 || preSelect == 2 {
//
//						} else {
//							continue
//						}
//						if wayRoute[currentBase][j] == 2 && preSelect != 2 {
//							wayLength[j] = wayLength[currentBase] + 1
//							consistentArrive[j] = -2 ^ preSelect
//						} else {
//							wayLength[j] = wayLength[currentBase] + 1
//							consistentArrive[j] = wayRoute[currentBase][j]
//						}
//						if result[j] != 0 {
//							result[j] = int(math.Min(float64(wayLength[j]), float64(result[j])))
//						} else {
//							result[j] = wayLength[j]
//						}
//					}
//				}
//				delete(consistentArrive, currentBase)
//			}
//		}
//	}
//	for i := 1; i < len(result); i++ {
//		if result[i] == 0 {
//			result[i] = -1
//		}
//	}
//	return result
//}

//func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
//	wayRoute := make([][]int, n)
//	var result = make([]int, n)
//	for i := 0; i < len(wayRoute); i++ {
//		wayRoute[i] = make([]int, n)
//	}
//	for i := 0; i < len(redEdges); i++ {
//		var x = redEdges[i][0]
//		var y = redEdges[i][1]
//		wayRoute[x][y] = 1
//	}
//	for i := 0; i < len(blueEdges); i++ {
//		var x = blueEdges[i][0]
//		var y = blueEdges[i][1]
//		if wayRoute[x][y] != 1 {
//			wayRoute[x][y] = -1
//		} else {
//			wayRoute[x][y] = 2
//		}
//	}
//
//	for i := 0; i < n; i++ {
//		if wayRoute[0][i] != 0 {
//			result[i] = 1
//			wayArrive := make(map[int]int)
//			wayArrive[i] = 1
//			preSelect := make(map[int]int)
//			preSelect[i] = wayRoute[0][i]
//			for len(wayArrive) > 0 {
//				var currentBase, step int
//				// random select on route
//				for key, value := range wayArrive {
//					currentBase = key
//					step = value
//					break
//				}
//				for j := 0; j < n; j++ {
//					if j == currentBase {
//						continue
//					}
//					if wayRoute[currentBase][j]^preSelect[currentBase] == -2 {
//						preSelect[j] = wayRoute[currentBase][j]
//					}
//				}
//			}
//		}
//	}
//
//}
