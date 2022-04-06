package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return int整型
 */
func FirstNotRepeatingChar(str string) int {
	// write code here
	var hash = make(map[rune]int)
	for _, v := range str {
		hash[v]++
	}
	for i, v := range str {
		if hash[v] == 1 {
			return i
		}
	}
	return -1
}
