package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
//递归,
func Mirror( pRoot *TreeNode ) *TreeNode {
    // write code here
    if pRoot == nil {
        return pRoot
    }

    Mirror(pRoot.Left)
    Mirror(pRoot.Right)

    pRoot.Left, pRoot.Right = pRoot.Right, pRoot.Left

    return pRoot
}
*/

//迭代
func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return pRoot
	}

	queue := []*TreeNode{pRoot}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		node.Left, node.Right = node.Right, node.Left

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return pRoot
}
