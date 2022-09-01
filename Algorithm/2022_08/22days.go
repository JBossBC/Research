package main

import (
	"math"
	"strconv"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func printTree(root *TreeNode) [][]string {
	var maxHeight = searchMaxDepth(root, 1)
	var result = make([][]string, maxHeight)
	var row = int(math.Pow(2, float64(maxHeight))) - 1
	for i := 0; i < len(result); i++ {
		result[i] = make([]string, row)
	}
	traversalTree(root, result, 1, (row-1)/2, maxHeight)
	return result
}
func searchMaxDepth(root *TreeNode, nowDepth int) int {
	if root == nil {
		return nowDepth - 1
	}
	return int(math.Max(float64(searchMaxDepth(root.Left, nowDepth+1)), float64(searchMaxDepth(root.Right, nowDepth+1))))
}

func traversalTree(root *TreeNode, elements [][]string, nowDepth int, nowLocation int, height int) {
	if root == nil {
		return
	}
	if nowDepth == 1 {
		elements[0][nowLocation] = strconv.FormatInt(int64(root.Val), 10)
	} else {
		elements[nowDepth-1][nowLocation] = strconv.FormatInt(int64(root.Val), 10)
	}
	traversalTree(root.Left, elements, nowDepth+1, nowLocation-int(math.Pow(2, float64(height-nowDepth-1))), height)
	traversalTree(root.Right, elements, nowDepth+1, nowLocation+int(math.Pow(2, float64(height-nowDepth-1))), height)
}
