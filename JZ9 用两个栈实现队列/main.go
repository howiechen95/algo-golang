package main

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 { //stack2为空时，移动stack1的内容到stack2
		for i := 0; i < len(stack1); i++ { //将stack1的数据，从后往前，全部移到stack2中
			stack2 = append(stack2, stack1[len(stack1)-1-i])
		}
		stack1 = []int{} //stack1置空
	}

	res := stack2[len(stack2)-1]    //获取最后一个位置作为栈顶元素
	stack2 = stack2[:len(stack2)-1] //左闭右开，取得0~n-2的元素作为新的切片
	return res                      //stack2弹出结果
}
