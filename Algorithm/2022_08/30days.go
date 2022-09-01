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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var arr = *mergeSort(root)
	arr = append(arr, val)
	return mergeCreate(arr, 0, len(arr)-1)
}

func mergeCreate(arr []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	var maxIndex = l
	for i := l; i <= r; i++ {
		if arr[maxIndex] < arr[i] {
			maxIndex = i
		}
	}
	var root = &TreeNode{
		Val:   arr[maxIndex],
		Left:  mergeCreate(arr, l, maxIndex-1),
		Right: mergeCreate(arr, maxIndex+1, r),
	}
	return root
}

func mergeSort(root *TreeNode) *[]int {
	if root == nil {
		return nil
	}

	var left = mergeSort(root.Left)
	var right = mergeSort(root.Right)
	var result []int
	if left != nil {
		result = append(result, *left...)
	}
	result = append(result, root.Val)
	if right != nil {
		result = append(result, *right...)
	}

	return &result
}
