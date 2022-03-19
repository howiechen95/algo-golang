package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @param pattern string字符串
 * @return bool布尔型
 */
func match(str string, pattern string) bool {
	n, m := len(str), len(pattern)
	str = " " + str
	pattern = " " + pattern
	s := []rune(str)
	p := []rune(pattern)
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i-1 >= 0 && p[j] != '*' {
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' {
				f[i][j] = (j-2 >= 0 && f[i][j-2]) || (i-1 >= 0 && f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.'))
			}
		}
	}

	return f[n][m]
}

func main() {
	fmt.Println(match("a", ".*"))
}
