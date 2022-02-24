package main

var S []int
var MinS []int // 单调递减栈
func Push(node int) {
	S = append(S, node)
	// 如果单调递减栈中无元素，或者栈顶元素大于等于插入元素，就入到单调栈中
	if len(MinS) == 0 || node <= MinS[len(MinS)-1] {
		MinS = append(MinS, node)
	}
}
func Pop() {
	if len(S) < 1 {
		return
	}
	tm := S[len(S)-1]
	S = S[:len(S)-1]
	// 如果出栈元素和单调栈元素相等，也出栈
	if tm == MinS[len(MinS)-1] {
		MinS = MinS[:len(MinS)-1]
	}
}
func Top() int {
	if len(S) < 1 {
		return -1
	}
	return S[len(S)-1]
}
func Min() int {
	if len(MinS) == 0 {
		return -1
	}
	// 返回单调栈栈顶元素
	return MinS[len(MinS)-1]
}
