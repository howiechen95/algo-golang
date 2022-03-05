package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var index = -1

func Serialize(root *TreeNode) string {
	result := ""
	if root == nil {
		result += "#!"
	} else {
		//result = result + strconv.Itoa(root.Val) + "!"
		//sleft := Serialize(root.Left)
		//sright := Serialize(root.Right)
		//result = result + sleft + sright
		result = result + strconv.Itoa(root.Val) + "!" + Serialize(root.Left) + Serialize(root.Right)
	}
	return result
}

func deseri(strs []string) *TreeNode {
	index++
	if strs[index] == "#" {
		return nil
	}
	val, _ := strconv.ParseInt(strs[index], 10, 64)
	root := &TreeNode{
		Val: int(val),
	}
	root.Left = deseri(strs)
	root.Right = deseri(strs)
	return root
}

func Deserialize(s string) *TreeNode {
	if s == "#!" || s == "" {
		return nil
	}
	ans := strings.Split(s, "!")
	return deseri(ans)
}
