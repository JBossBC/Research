package main

import (
	"reflect"
	"sort"
	"unsafe"
)

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	// find head
	var convertArr = make([]int, 0)
	var temp = head
	for temp != nil {
		convertArr = append(convertArr, temp.Val)
		temp = temp.Next
	}
	reverse := sort.Reverse(sort.IntSlice(convertArr))
	pointer := *unsafe.Pointer(reflect.ValueOf(reverse).Pointer())
	slice := unsafe.Slice(&pointer, len(convertArr))

}
