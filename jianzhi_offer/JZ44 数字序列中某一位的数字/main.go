package main

import (
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param n int整型
 * @return int整型
 */
func findNthDigit(n int) int {
	if n <= 0 {
		return 0
	}
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = start * 9 * digit
	}
	num := strconv.Itoa(start + (n-1)/digit)
	idx := (n - 1) % digit
	var s string = num[idx : idx+1]
	r, _ := strconv.ParseInt(s, 10, 64)
	return int(r)
}

func main() {
	findNthDigit(1)
}
