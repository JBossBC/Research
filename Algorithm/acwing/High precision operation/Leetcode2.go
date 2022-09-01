package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//动态规划
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var i, j = l1, l2
	var temp int
	for i != nil && j != nil {
		if i != nil {
			temp += i.Val
		} else {
			i.Next = &ListNode{Val: 0}
		}
		if j != nil {
			temp += j.Val
		}
		i.Val = temp % 10
		i = i.Next
		j = j.Next
		temp /= 10
	}

	return l1
}
