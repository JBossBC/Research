package main

func main() {

}
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	var temp = root.Right
	root.Right = root.Left
	root.Left = temp
	return root
}
