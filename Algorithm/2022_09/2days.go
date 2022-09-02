package main

import (
	"fmt"
	"math/rand"
)

func main() {
	AutoCompareDevice()
}
func AutoCompareDevice() {
	//var before = time.Now()
	rand.Seed(500000)
	for i := 0; i < 1000000; i++ {
		var arr = rand.Perm(rand.Intn(3000))
		var root = createTree(arr, 1, nil)
		var myAnswer = longestUnivaluePath(root)
		fmt.Println(myAnswer)
	}
	//var standard = widthOfBinaryTree(root)
	//	if myAnswer != standard {
	//		log.Println("method is wrong,please update")
	//		fmt.Printf("myAnswer is %d,the standard is %d\n", myAnswer, standard)
	//		var temp = root
	//		recursionPrintTree(temp)
	//		return
	//	}
	//	resultMap = make(map[int]*LengthLocation, 3000)
	//}
	//var after = time.Now()
	//time.Ticker{}
	//fmt.Println(time.NewTimer(before))
	fmt.Println("success")
}
func createTree(nums []int, NodeNumber int, node *TreeNode) *TreeNode {
	var temp *TreeNode
	if NodeNumber-1 >= len(nums) {
		return nil
	}
	var randNumber = rand.Intn(100)
	if nums[NodeNumber-1] > randNumber {
		temp = &TreeNode{
			Val: nums[NodeNumber-1],
		}
	}
	if temp == nil {
		return nil
	}
	if NodeNumber*2-1 < len(nums) {
		temp.Left = &TreeNode{Val: nums[NodeNumber*2-1]}
		createTree(nums, NodeNumber*2, temp.Left)
	}
	if NodeNumber*2 < len(nums) {
		temp.Right = &TreeNode{Val: nums[NodeNumber*2]}
		createTree(nums, NodeNumber*2+1, temp.Right)
	}
	return temp
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memory map[*TreeNode]int

func longestUnivaluePath(root *TreeNode) int {
	memory = make(map[*TreeNode]int, 10000)
	var max = 0
	for _, value := range memory {
		if value > max {
			max = value
		}
	}
	return max
}
func recursion(root *TreeNode, preRoot *TreeNode) {
	if root == nil {
		return
	}
	if root.Val == preRoot.Val {
		memory[preRoot]++
		recursion(root.Left, preRoot)
		recursion(root.Right, preRoot)
	}
	recursion(root.Left, root)
	recursion(root.Right, root)
}
