package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param n int整型
 * @param m int整型
 * @return int整型
 */
func LastRemaining_Solution(n int, m int) int {
	// write code here
	if n <= 0 {
		return -1
	}
	return f(n, m)
}

func f(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (f(n-1, m) + m) % n
}
