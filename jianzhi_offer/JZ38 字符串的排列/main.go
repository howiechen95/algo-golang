package main

import "sort"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
func Permutation(str string) []string {
	// write code here
	// write code here
	n := len(str)
	if n == 0 {
		return nil
	}

	// 转换为切片是为了 排序
	bs := []byte(str)
	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	res := make([]string, 0)
	used := make([]bool, n)
	path := make([]byte, 0)

	var do func(path []byte)
	do = func(path []byte) {
		// 长度复合，处理值并 返回
		if len(path) == n {
			s := string(path)
			res = append(res, s)
			return
		}

		for i := 0; i < n; i++ {
			// 去重
			if used[i] || i > 0 && bs[i] == bs[i-1] && !used[i-1] {
				continue
			}

			// 回溯
			used[i] = true
			path = append(path, bs[i])

			do(path)

			used[i] = false
			path = path[:len(path)-1]
		}
	}

	do(path)

	return res
}
