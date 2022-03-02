package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	return IsSame(pRoot, pRoot)
}

func IsSame(p1 *TreeNode, p2 *TreeNode) bool {
	if p1 == nil && p2 == nil {
		return true
	}
	if (p1 == nil && p2 != nil) || (p1 != nil && p2 == nil) {
		return false
	}
	return p1.Val == p2.Val && IsSame(p1.Left, p2.Right) && IsSame(p1.Right, p2.Left)
}
