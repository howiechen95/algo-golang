package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param A int整型一维数组
 * @return int整型一维数组
 */
func multiply(A []int) []int {
	// write code here
	if len(A) == 0 {
		return A
	}
	res := make([]int, len(A))
	res[0] = 1
	for i := 1; i < len(A); i++ {
		res[i] = res[i-1] * A[i-1]
	}
	tmp := 1
	for i := len(A) - 2; i >= 0; i-- {
		tmp *= A[i+1]
		res[i] *= tmp
	}
	return res
}
