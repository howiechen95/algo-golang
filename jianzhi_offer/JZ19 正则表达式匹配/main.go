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
	// 技巧：往原字符头部插入空格，这样得到 char 数组是从 1 开始，而且可以使得 f[0][0] = true，可以将 true 这个结果滚动下去
	n, m := len(str), len(pattern)
	str = " " + str
	pattern = " " + pattern
	s := []rune(str)
	p := []rune(pattern)
	// f(i,j) 代表考虑 s 中的 1~i 字符和 p 中的 1~j 字符 是否匹配
	f := make([][]bool, n+1)
	for i := range f {
		f[i] = make([]bool, m+1)
	}

	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 如果下一个字符是 '*'，则代表当前字符不能被单独使用，跳过
			if j+1 <= m && p[j+1] == '*' {
				continue
			}
			if i-1 >= 0 && p[j] != '*' { // 对应了 p[j] 为普通字符和 '.' 的两种情况
				f[i][j] = f[i-1][j-1] && (s[i] == p[j] || p[j] == '.')
			} else if p[j] == '*' { // 对应了 p[j] 为 '*' 的情况
				f[i][j] = (j-2 >= 0 && f[i][j-2]) || (i-1 >= 0 && f[i-1][j] && (s[i] == p[j-1] || p[j-1] == '.'))
			}
		}
	}

	return f[n][m]
}

func main() {
	fmt.Println(match("a", ".*"))
}
