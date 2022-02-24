package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param proot TreeNode类
 * @param k int整型
 * @return int整型
 */
func KthNode(proot *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for proot != nil {
			stack = append(stack, proot)
			proot = proot.Left
		}
		if len(stack) == 0 {
			return -1
		}
		stack, proot = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return proot.Val
		}
		proot = proot.Right
	}
}
