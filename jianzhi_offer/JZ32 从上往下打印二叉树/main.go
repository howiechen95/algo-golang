package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PrintFromTopToBottom(root *TreeNode) []int {
	var ans []int
	if root == nil { //空树
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return ans
}
