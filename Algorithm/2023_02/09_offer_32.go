package main

func main() {
	t := &TreeNode{Val: 3}
	levelOrder(t)
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

func levelOrder(root *TreeNode) [][]int {
	var queue = make([]*TreeNode, 1010)
	var head, tail int = 0, 0
	var nextQueue = make([]*TreeNode, 1010)
	var nextHead, nextTail int = 0, -1
	queue[0] = root
	var result = make([][]int, 0, 500)
	var height = 0
	var currentArray = make([]int, 0, 500)
	for head <= tail {
		var temp = queue[head]
		head++
		if temp == nil {
			break
		}
		currentArray = append(currentArray, temp.Val)
		if temp.Left != nil {
			nextTail = nextTail + 1
			nextQueue[nextTail] = temp.Left
		}
		if temp.Right != nil {
			nextTail = nextTail + 1
			nextQueue[nextTail] = temp.Right
		}
		//元素被消耗完,尝试下一个高度
		if head > tail {
			var temp = nextQueue
			nextQueue = queue
			queue = temp
			head = nextHead
			tail = nextTail
			height++
			nextHead = 0
			nextTail = -1
			result = append(result, currentArray)
			currentArray = make([]int, 0, 500)
		}
	}
	for i := 0; i < height; i++ {
		if i%2 == 0 {
			continue
		}
		var left, right = 0, len(result[i]) - 1
		var lineArr = result[i]
		for left < right {
			if lineArr[left] == lineArr[right] {
				left++
				right--
				continue
			}
			lineArr[left] = lineArr[left] ^ lineArr[right]
			lineArr[right] = lineArr[left] ^ lineArr[right]
			lineArr[left] = lineArr[left] ^ lineArr[right]
			left++
			right--
		}
	}
	return result[:height]
}
