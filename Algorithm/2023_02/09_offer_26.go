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
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	var queue = make([]*TreeNode, 10010)
	var head, tail int = 0, 0
	queue[0] = A
	var nextQueue = make([]*TreeNode, 10010)
	var nextHead, nextTail = 0, -1
	for head <= tail {
		var temp = queue[head]
		head++
		if temp == nil {
			break
		}
		if temp.Val == B.Val {
			result := judge(temp, B)
			if result == 1 {
				return true
			}
		}
		if temp.Left != nil {
			nextTail = nextTail + 1
			nextQueue[nextTail] = temp.Left
		}
		if temp.Right != nil {
			nextTail = nextTail + 1
			nextQueue[nextTail] = temp.Right
		}
		if head > tail {
			var temp = queue
			queue = nextQueue
			nextQueue = temp
			head = nextHead
			tail = nextTail
			nextHead = 0
			nextTail = -1
		}
	}
	return false
}
func judge(A *TreeNode, B *TreeNode) int {
	if B == nil {
		return 1
	}
	if A == nil {
		return 0
	}
	var result int
	if A.Val == B.Val {
		result = 1
	} else {
		return 0
	}
	var leftResult, rightResult int = 1, 1
	if B.Left != nil {
		if A.Left == nil {
			return 0
		}
		leftResult = judge(A.Left, B.Left)
	}
	if B.Right != nil {
		if A.Right == nil {
			return 0
		}
		rightResult = judge(A.Right, B.Right)
	}
	return leftResult & result & rightResult
}
