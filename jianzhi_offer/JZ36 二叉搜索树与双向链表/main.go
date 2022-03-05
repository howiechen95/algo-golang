package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var pre, root *TreeNode

// 空间复杂度 O(1) 解法
func Convert(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return pRootOfTree
	}
	// 递归左子树
	Convert(pRootOfTree.Left)
	// 判断特殊情况
	if root == nil {
		root = pRootOfTree
	}
	// 修改遍历结点为双向链表
	if pre != nil {
		pRootOfTree.Left = pre
		pre.Right = pRootOfTree
	}
	// 更新 pre
	pre = pRootOfTree
	// 递归右子树
	Convert(pRootOfTree.Right)
	return root
}

var treeList []*TreeNode

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	treeList = append(treeList, root)
	inorder(root.Right)
}

func Convert2(pRootOfTree *TreeNode) *TreeNode {
	if pRootOfTree == nil {
		return pRootOfTree
	}
	inorder(pRootOfTree)
	for i := 0; i < len(treeList)-1; i++ {
		treeList[i].Right = treeList[i+1]
		treeList[i+1].Left = treeList[i]
	}
	return treeList[0]
}
