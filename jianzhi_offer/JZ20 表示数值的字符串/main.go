package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @return bool布尔型
 */
import (
	"strconv"
	"strings"
)

func isXiaoshu(str string) bool {
	if _, err := strconv.ParseFloat(strings.Trim(str, " "), 64); err != nil {
		return false
	} else {
		return true
	}
}
func isZhengshu(str string) bool {
	if _, err := strconv.Atoi(strings.Trim(str, " ")); err != nil {
		return false
	} else {
		return true
	}
}
func isKexue(str string) bool {
	if strings.Count(str, "e") == 1 {
		strs := strings.Split(str, "e")
		return (isXiaoshu(strs[0]) || isZhengshu(strs[0])) && isZhengshu(strs[1])
	} else if strings.Count(str, "E") == 1 {
		strs := strings.Split(str, "E")
		return (isXiaoshu(strs[0]) || isZhengshu(strs[0])) && isZhengshu(strs[1])
	} else {
		return false
	}
}
func isNumeric(str string) bool {
	// write code here
	return isXiaoshu(str) || isZhengshu(str) || isKexue(str)
}
