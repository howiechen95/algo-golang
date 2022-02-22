package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return [][]int{}
	}

	queue := []*TreeNode{pRoot}
	levels := [][]int{}

	for len(queue) > 0 {
		n := len(queue)
		level := []int{}

		for i := 0; i < n; i++ {
			pRoot = queue[0]
			queue = queue[1:]

			level = append(level, pRoot.Val)
			if pRoot.Left != nil {
				queue = append(queue, pRoot.Left)
			}
			if pRoot.Right != nil {
				queue = append(queue, pRoot.Right)
			}
		}

		k1 := len(level)  //这里注意，内层【】，翻转的都是这里面的
		k2 := len(levels) //这里也要注意，外层【】，只有判断奇偶的时候用
		if k2%2 == 1 {
			for i := 0; i < k1/2; i++ {
				level[i], level[k1-1-i] = level[k1-1-i], level[i]
			}
		}
		levels = append(levels, level)
	}
	return levels
}
