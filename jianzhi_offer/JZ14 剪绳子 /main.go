package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func cutRope(number int) int {
	// write code here
	if number%3 == 0 {
		return power(3, number/3)
	}
	if number%3 == 1 {
		return power(3, (number-1)/3-1) * 4
	}
	if number%3 == 2 {
		return power(3, (number-2)/3) * 2
	}
	return 0
}

func power(num int, n int) int {
	if n == 0 {
		return 1
	} else {
		return num * power(num, n-1)
	}
}
