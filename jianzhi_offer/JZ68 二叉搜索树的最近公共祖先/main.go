package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	curNode := root
	for {
		if p < curNode.Val && q < curNode.Val {
			curNode = curNode.Left
		} else if p > curNode.Val && q > curNode.Val {
			curNode = curNode.Right
		} else {
			return curNode.Val
		}
	}
}
