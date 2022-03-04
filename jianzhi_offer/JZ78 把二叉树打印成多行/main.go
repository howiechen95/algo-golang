package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	var result [][]int
	if pRoot == nil {
		return result
	}

	// write code here
	queue := list.New()
	queue.PushBack(pRoot)

	for queue.Len() > 0 {
		curlen := queue.Len()
		var curList []int
		for i := 0; i < curlen; i++ {
			curTree := queue.Remove(queue.Front()).(*TreeNode)
			curList = append(curList, curTree.Val)
			if curTree.Left != nil {
				queue.PushBack(curTree.Left)
			}

			if curTree.Right != nil {
				queue.PushBack(curTree.Right)
			}

		}

		result = append(result, curList)
	}

	return result

}
