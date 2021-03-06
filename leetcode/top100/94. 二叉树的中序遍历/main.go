package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	left = append(left, root.Val)
	for _, v := range right {
		left = append(left, v)
	}
	return left
}
