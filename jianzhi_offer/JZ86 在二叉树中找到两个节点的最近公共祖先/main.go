package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	if root == nil {
		return 0
	}
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}
	Left := lowestCommonAncestor(root.Left, o1, o2)
	Right := lowestCommonAncestor(root.Right, o1, o2)
	if Left != 0 && Right != 0 {
		return root.Val
	}
	if Left != 0 {
		return Left
	}
	return Right
}
