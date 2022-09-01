package main

import "fmt"

func main() {

	deepestLeavesSum(nil)

}

/**
  Definition for a binary tree node.
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

var result = make(map[int]int, 0)
var maxDepth = 0

/**
trouble:tree combined by linklist which node is hard to find
*/
func deepestLeavesSum(root *TreeNode) int {
	oneLeft := TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}
	oneRight := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	TestRoot := TreeNode{
		Val:   6,
		Left:  &oneLeft,
		Right: &oneRight,
	}
	Digui(&TestRoot, 0)
	fmt.Println(result[maxDepth])
	return result[maxDepth]
}
func Digui(temp *TreeNode, depth int) {
	if temp == nil {
		return
	}
	if temp.Left == nil && temp.Right == nil {
		if depth > maxDepth {
			maxDepth = depth
		}
		result[depth] = result[depth] + temp.Val
		return
	}
	Digui(temp.Left, depth+1)
	Digui(temp.Right, depth+1)
}
