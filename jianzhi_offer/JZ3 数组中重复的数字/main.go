package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	count := make([]int, 10001)
	for _, n := range numbers {
		if count[n] > 0 {
			return n
		}
		count[n]++
	}
	return -1
}
