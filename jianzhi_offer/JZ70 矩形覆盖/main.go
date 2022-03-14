package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param number int整型
 * @return int整型
 */
func rectCover(number int) int {
	if number <= 2 {
		return number
	}
	pre1, pre2 := 2, 1
	for i := 3; i <= number; i++ {
		cur := pre1 + pre2
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}
