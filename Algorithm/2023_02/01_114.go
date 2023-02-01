package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func flatten(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	LeftAction(root)
//}
//
//func LeftAction(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	LeftAction(root.Left)
//	LeftAction(root.Right)
//	if root.Left == nil {
//		return
//	}
//	var tail = root.Right
//	//swap
//	root.Right = root.Left
//	root.Left = nil
//	var endPoint = root.Right
//	for endPoint.Right != nil {
//		endPoint = endPoint.Right
//	}
//	endPoint.Right = tail
//}
