package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func Convert(pRootOfTree *TreeNode) *TreeNode {
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
