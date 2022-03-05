package main

type TreeLinkNode struct {
	Val   int
	Left  *TreeLinkNode
	Right *TreeLinkNode
	Next  *TreeLinkNode
}

var treeList []*TreeLinkNode

func inorder(root *TreeLinkNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	treeList = append(treeList, root)
	inorder(root.Right)
}
func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	var root *TreeLinkNode
	tmp := pNode
	for tmp != nil {
		root = tmp
		tmp = tmp.Next
	}
	inorder(root)

	for i, n := 0, len(treeList); i < n; i++ {
		if treeList[i] == pNode && i+1 != n {
			return treeList[i+1]
		}
	}
	return nil
}
