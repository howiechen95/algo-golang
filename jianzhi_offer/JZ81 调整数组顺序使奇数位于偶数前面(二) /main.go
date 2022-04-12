package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func reOrderArrayTwo(array []int) []int {
	// write code here
	var i, j int
	for i = 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			for j = i + 1; j < len(array); j++ {
				if array[j]%2 != 0 {
					array[i], array[j] = array[j], array[i]
					break
				}
			}
		}
	}
	return array
}
