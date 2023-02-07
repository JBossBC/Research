package main

func main() {
	var right1Node = &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	var root = &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9},
		Right: right1Node,
	}

	levelOrder(root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var result []int
var curPointer []*TreeNode
var nextPointer []*TreeNode

func levelOrder(root *TreeNode) []int {
	result = make([]int, 0, 1000)
	nextPointer = make([]*TreeNode, 500)
	curPointer = make([]*TreeNode, 500)
	var nextIndex = -1
	if root == nil {
		return nil
	}
	curPointer[0] = root
	var currentIndex = 0
	for i := 0; i <= currentIndex; i++ {
		var currentNode = curPointer[i]
		result = append(result, currentNode.Val)
		if currentNode.Left != nil {
			nextIndex++
			nextPointer[nextIndex] = currentNode.Left
		}
		if currentNode.Right != nil {
			nextIndex++
			nextPointer[nextIndex] = currentNode.Right
		}
		if i == currentIndex {
			var temp = curPointer
			curPointer = nextPointer
			nextPointer = temp
			currentIndex = nextIndex
			nextIndex = -1
			i = -1
		}
	}
	return result
}
