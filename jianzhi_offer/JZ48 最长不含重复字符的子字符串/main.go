package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return int整型
 */
func lengthOfLongestSubstring(s string) int {
	count := make(map[string]int)
	maxLength := 0
	r := 0
	n := len(s)
	for l := 0; l < n; l++ {
		for r < n {
			if _, ok := count[string(s[r])]; !ok {
				count[string(s[r])] = 1
				r += 1
			} else {
				break
			}
		}
		if r-l > maxLength {
			maxLength = r - l
		}
		count[string(s[l])] -= 1
		if count[string(s[l])] == 0 {
			delete(count, string(s[l]))
		}
	}
	return maxLength
}
