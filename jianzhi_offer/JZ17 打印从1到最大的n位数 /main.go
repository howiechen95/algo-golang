package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型 最大位数
 * @return int整型一维数组
 */
func printNumbers(n int) []int {
	// write code here
	end := 1
	x := 10
	for n != 0 { //这波叫学以致用，快速幂
		if (n & 1) == 1 {
			end = end * x
		}
		x = x * x
		n = n / 2
	}
	res := []int{}
	for i := 1; i < end; i++ {
		res = append(res, i)
	}
	return res
}
