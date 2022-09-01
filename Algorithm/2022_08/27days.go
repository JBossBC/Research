package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

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

func main() {
	// var left = &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 5,
	// 	},
	// }
	// root := &TreeNode{
	// 	Val:  1,
	// 	Left: left,
	// 	Right: &TreeNode{
	// 		Val: 2,
	// 	},
	// }
	// var result = widthOfBinaryTree(root)
	// fmt.Println(result)
	AutoCompareDevice()
}

type LengthLocation struct {
	Left  int
	Right int
}

var resultMap map[int]*LengthLocation = make(map[int]*LengthLocation, 3000)

func myWidthOfBinaryTree(root *TreeNode) int {
	recursion(root, 1, 1)
	var max = 0
	for _, value := range resultMap {
		if value.Left == 0 || value.Right == 0 {
			continue
		}
		var temp = value.Right - value.Left + 1
		if temp > max {
			max = temp
		}
	}
	if max < 1 && root != nil {
		max = 1
	}
	return max
}
func recursion(root *TreeNode, rootNode int, length int) {
	if root == nil {
		return
	}
	value, ok := resultMap[length]
	if !ok {
		//这里能保证为分配的是比较晓得,通过遍历顺序可以明确
		resultMap[length] = &LengthLocation{
			Left: rootNode,
		}
	} else {
		if rootNode > value.Right {
			value.Right = rootNode
		}

	}
	recursion(root.Left, 2*rootNode, length+1)
	recursion(root.Right, 2*rootNode+1, length+1)
}

func AutoCompareDevice() {
	var before = time.Now()
	rand.Seed(500000)
	for i := 0; i < 1000000; i++ {
		var arr = rand.Perm(rand.Intn(3000))
		var root = createTree(arr, 1, nil)
		var myAnswer = myWidthOfBinaryTree(root)
		var standard = widthOfBinaryTree(root)
		if myAnswer != standard {
			log.Println("method is wrong,please update")
			fmt.Printf("myAnswer is %d,the standard is %d\n", myAnswer, standard)
			var temp = root
			recursionPrintTree(temp)
			return
		}
		resultMap = make(map[int]*LengthLocation, 3000)
	}
	var after = time.Now()
	time.Ticker{}
	fmt.Println(time.NewTimer(before))
	fmt.Println("success")
}
func recursionPrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%4d", root.Val)
	recursionPrintTree(root.Left)
	recursionPrintTree(root.Right)
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

func widthOfBinaryTree(root *TreeNode) int {
	levelMin := map[int]int{}
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, depth, index int) int {
		if node == nil {
			return 0
		}
		if _, ok := levelMin[depth]; !ok {
			levelMin[depth] = index // 每一层最先访问到的节点会是最左边的节点，即每一层编号的最小值
		}
		return max(index-levelMin[depth]+1, max(dfs(node.Left, depth+1, index*2), dfs(node.Right, depth+1, index*2+1)))
	}
	return dfs(root, 1, 1)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
