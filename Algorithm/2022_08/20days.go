package main

import (
	"fmt"
	"math"
)

func main() {
	tree := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(tree)
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

//理解出错
//func constructMaximumBinaryTree(nums []int) *TreeNode {
//	if len(nums) <= 0 {
//		return nil
//	}
//	var maxIndex = -1
//	var maxValue = math.MinInt
//	for index, value := range nums {
//		if maxValue < value {
//			maxValue = value
//			maxIndex = index
//		}
//	}
//	var root = &TreeNode{
//		Val: maxValue,
//	}
//	if maxIndex != 0 {
//		var leftArr = make([]int, len(nums[:maxIndex]))
//		copy(leftArr, nums[:maxIndex])
//		root.Left = &TreeNode{
//			Val: nums[0],
//		}
//		createTreeNode(root.Left, leftArr, 1, 1)
//	}
//	//add 1 maybe appear error(array index out of len), select rightArr add maxValue to avoid error
//	if maxIndex != len(nums)-1 {
//		var rightArr = make([]int, len(nums[maxIndex+1:len(nums)]))
//		copy(rightArr, nums[maxIndex+1:len(nums)])
//		root.Right = &TreeNode{
//			Val: len(nums) - 1,
//		}
//		createTreeNode(root.Right, rightArr, len(rightArr)-2, -1)
//	}
//	return root
//}
//func createTreeNode(root *TreeNode, addValueArr []int, needAddIndex int, nextOperation int) {
//	if needAddIndex >= len(addValueArr) || needAddIndex < 0 {
//		return
//	}
//	if addValueArr[needAddIndex] > root.Val {
//		root.Right = &TreeNode{
//			Val: addValueArr[needAddIndex],
//		}
//		createTreeNode(root.Right, addValueArr, needAddIndex+nextOperation, nextOperation)
//	}
//	if addValueArr[needAddIndex] < root.Val {
//		root.Left = &TreeNode{
//			Val: addValueArr[needAddIndex],
//		}
//		createTreeNode(root.Left, addValueArr, needAddIndex+nextOperation, nextOperation)
//	}
//}
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return recursionAdd(nums, nil)
}
func recursionAdd(nums []int, root *TreeNode) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	var maxIndex = -1
	var maxValue = math.MinInt
	for index, value := range nums {
		if maxValue < value {
			maxIndex = index
			maxValue = value
		}
	}
	root = &TreeNode{
		Val: maxValue,
	}
	left := recursionAdd(nums[:maxIndex], root.Left)
	root.Left = left
	if maxIndex+1 < len(nums) {
		right := recursionAdd(nums[maxIndex+1:], root.Right)
		root.Right = right
	}
	return root
}
