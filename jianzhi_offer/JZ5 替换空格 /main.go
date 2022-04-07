package main

import "strings"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return string字符串
 */
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
