package main

func main() {
  &TreeNode{Val: 1,Left: &TreeNode{Val: 2,Left: TreeNode{Val: }}}
}

func isSymmetric(root *TreeNode) bool {
	if root  == nil{
		return true
	}
	if root.Left ==nil || root.Right ==nil{
		if root.Left==nil && root.Right==nil {
			goto begin
		}
	}
begin:
	var result=judgeEqual(root.Left,root.Right)
	if result ==1{
		return true
	}
	return false

}

func judgeEqual(left *TreeNode,right *TreeNode)int{
	if left == nil &&right ==nil{
		return 1
	}
	if left ==nil{
		return 0
	}
	if right == nil{
		return  0
	}
	var result =0
	if left.Val ==right.Val{
		result =1
	}
	if result == 0{
		return 0
	}
	var leftResult,rightResult int =1,1
	if left.Left==nil || right.Right ==nil{
		if left.Left ==nil&&right.Right==nil{
			leftResult=1
		}else{
			return 0
		}
	}else{
		leftResult=judgeEqual(left.Left,right.Right)
	}
	if left.Right==nil || right.Left ==nil{
		if left.Right ==nil&&right.Left==nil{
			rightResult=1
		}else{
			return 0
		}
	}else{
		rightResult=judgeEqual(left.Right,right.Left)
	}
	return result&leftResult&rightResult
}