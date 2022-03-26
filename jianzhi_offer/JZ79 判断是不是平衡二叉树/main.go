package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//自上而下递归， 最坏时间On^2，平均时间Onlogn, 空间on
func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	return abs(height(pRoot.Left)-height(pRoot.Right)) <= 1 && IsBalanced_Solution(pRoot.Left) && IsBalanced_Solution(pRoot.Right)
}

func height(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	return max(height(pRoot.Left), height(pRoot.Right)) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
