package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param sequence int整型一维数组
 * @return bool布尔型
 */
func VerifySquenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return false
	}

	stack := []int{}
	min, max := -1, 1<<31-1
	stack = append(stack, min)

	for i := len(sequence) - 1; i >= 0; i-- {
		if sequence[i] > max {
			return false
		}

		for sequence[i] < stack[len(stack)-1] {
			max = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, sequence[i])
	}

	return true
}
