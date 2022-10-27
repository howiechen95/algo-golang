package main

import "fmt"

func lengthOfLastWord(s string) int {
	l := 0
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}
	for ; i >= 0; i-- {
		if s[i] == ' ' {
			break
		} else {
			l++
		}
	}
	return l
}

func main() {
	testcases := []string{"", "aa bb", "aa bb ", "  ", "abc"}
	for _, v := range testcases {
		fmt.Println(lengthOfLastWord(v))
	}
}
