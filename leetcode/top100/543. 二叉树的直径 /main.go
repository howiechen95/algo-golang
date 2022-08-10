package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var lengthCom func(t *TreeNode) int
	var max_num int = math.MinInt64
	lengthCom = func(t *TreeNode) int {
		var l, r int
		if t.Left != nil {
			l = 1 + lengthCom(t.Left)
		} else {
			l = 0
		}
		if t.Right != nil {
			r = 1 + lengthCom(t.Right)
		} else {
			r = 0
		}
		max_num = max(max_num, l+r)
		return max(l, r)
	}
	lengthCom(root)
	return max_num
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
