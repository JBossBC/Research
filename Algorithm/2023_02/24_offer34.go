package main

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
var answer [][]int

func pathSum(root *TreeNode, target int) [][]int {
	answer = make([][]int, 0, 2500)
	result := make([]int, 0, 2500)
	dfs(root, result, target)
	return answer
}

func dfs(root *TreeNode, result []int, remain int) {
	if root == nil {
		return
	}
	remain -= root.Val
	result = append(result, root.Val)
	if root.Left == nil && root.Right == nil {
		if remain == 0 {
			var r []int = make([]int, len(result))
			copy(r, result)
			answer = append(answer, r)
		}
	}
	var currentLength = len(result)
	if root.Left != nil {
		dfs(root.Left, result, remain)
	}
	result = result[:currentLength]
	if root.Right != nil {
		dfs(root.Right, result, remain)
	}

}
