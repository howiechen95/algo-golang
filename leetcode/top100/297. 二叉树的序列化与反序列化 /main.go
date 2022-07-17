package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

func (this Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	left := "(" + this.serialize(root.Left) + ")"
	right := "(" + this.serialize(root.Right) + ")"
	return left + strconv.Itoa(root.Val) + right
}

func (Codec) deserialize(data string) *TreeNode {
	var parse func() *TreeNode
	parse = func() *TreeNode {
		if data[0] == 'X' {
			data = data[1:]
			return nil
		}
		node := &TreeNode{}
		data = data[1:] // 跳过左括号
		node.Left = parse()
		data = data[1:] // 跳过右括号
		i := 0
		for data[i] == '-' || '0' <= data[i] && data[i] <= '9' {
			i++
		}
		node.Val, _ = strconv.Atoi(data[:i])
		data = data[i:]
		data = data[1:] // 跳过左括号
		node.Right = parse()
		data = data[1:] // 跳过右括号
		return node
	}
	return parse()
}

func main() {
	Constructor()
}
