package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	leftDepth := TreeDepth(pRoot.Left)
	rightDepth := TreeDepth(pRoot.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}
