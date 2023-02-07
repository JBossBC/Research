package main

func main() {

}

/**
* Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//var mapping = map[int]interface{}{1: nil, 0: nil}
//
//func evaluateTree(root *TreeNode) bool {
//	var result = digui(root)
//	if result == 0 {
//		return false
//	} else {
//		return true
//	}
//
//}
//
//func digui(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var leftResult = digui(root.Left)
//	var rightResult = digui(root.Right)
//	if _, ok := mapping[root.Val]; ok {
//		return root.Val
//	}
//	if root.Val == 2 {
//		return leftResult | rightResult
//	} else {
//		return leftResult & rightResult
//	}
//}
