package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloorII(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	return 1 << uint(number-1)
}
