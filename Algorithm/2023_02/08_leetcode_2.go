package main

func main() {
	//var l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	//var l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	var l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	var l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	addTwoNumbers(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var root *ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var nextPointer = &ListNode{}
	root = nextPointer
	var carry = 0
	for l1 != nil || l2 != nil {
		var nextL1, nextL2 *ListNode
		var value1, value2 int
		if l1 == nil {
			value1 = 0
			nextL1 = nil
		} else {
			nextL1 = l1.Next
			value1 = l1.Val
		}
		if l2 == nil {
			value2 = 0
			nextL2 = nil
		} else {
			value2 = l2.Val
			nextL2 = l2.Next
		}
		nextPointer.Val = value1 + value2 + carry
		if nextPointer.Val >= 10 {
			nextPointer.Val = nextPointer.Val % 10
			carry = 1
		} else {
			carry = 0
		}
		l1 = nextL1
		l2 = nextL2
		if l1 != nil || l2 != nil {
			nextPointer.Next = &ListNode{}
			nextPointer = nextPointer.Next
		} else {
			if carry != 0 {
				nextPointer.Next = &ListNode{Val: 1}
			}
		}
	}
	return root
}
