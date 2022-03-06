package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var key int

func dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	// 注意这里为什么不return？
	// - 题目定义不需要从根节点开始，也不需要在叶子节点结束；
	// - 因此需要从根节点开始，一个节点一个节点的进行深度遍历；
	// - 也就是从根节点到叶子节点可能有多条可能的路径；
	// - 可能是无序的小于等于0的Val；
	// - 因此只要保证后继连续节点Value和为0，就又是一条新路径；
	if sum == 0 {
		key++
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}
func FindPath(root *TreeNode, sum int) int {
	if root == nil {
		return key
	}
	// 题目定义不需要从根节点开始，也不需要在叶子节点结束；
	// 因此需要从根节点开始，一个节点一个节点的进行深度遍历；
	dfs(root, sum)
	FindPath(root.Left, sum)
	FindPath(root.Right, sum)
	return key
}
