package main

import "strconv"

func NumberOf1Between1AndN_Solution(n int) int {
	a := 0
	for i := 1; i <= n; i++ {
		b := strconv.Itoa(i)
		for _, val := range []byte(b) {
			if val == '1' {
				a++
			}
		}
	}
	return a
}
