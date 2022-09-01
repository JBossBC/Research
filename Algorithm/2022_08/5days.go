package main

import "fmt"

func main() {
	var call TreeNode
	fmt.Println(call)
	var leftNode = &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
	}
	var root = &TreeNode{
		Val:  1,
		Left: leftNode,

		Right: &TreeNode{
			Val: 3,
		},
	}
	addOneRow(root, 5, 4)

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var memoryPre = []*TreeNode{root}
	var currentDepth = 1
	if depth == 1 {
		var newRoot = &TreeNode{
			Val:  val,
			Left: root,
		}
		return newRoot
	}
	for currentDepth != depth-1 {
		memoryPre = replaceMemoryPre(memoryPre)
		currentDepth++
	}
	for _, value := range memoryPre {

		var CommonLeftNode = &TreeNode{
			Val: val,
		}
		var commonRightNode = &TreeNode{
			Val: val,
		}
		if value.Left != nil {
			CommonLeftNode.Left = value.Left
		}
		if value.Right != nil {
			commonRightNode.Right = value.Right
		}
		value.Left = CommonLeftNode
		value.Right = commonRightNode
	}
	return root
}
func replaceMemoryPre(memoryPre []*TreeNode) []*TreeNode {
	var resultArr []*TreeNode
	for _, value := range memoryPre {
		var addArr []*TreeNode
		if value.Left != nil {
			addArr = append(addArr, value.Left)
		}
		if value.Right != nil {
			addArr = append(addArr, value.Right)
		}
		resultArr = append(resultArr, addArr...)
	}
	return resultArr
}
