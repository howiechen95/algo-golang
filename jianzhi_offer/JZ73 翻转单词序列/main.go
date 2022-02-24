package main

func ReverseSentence(str string) string {
	n := len(str)
	ans := ""
	l, r := 0, n-1
	for l < n && str[l] == ' ' {
		l++
	}
	for r >= 0 && str[r] == ' ' {
		r--
	}
	j := r + 1
	for i := r; i >= l; i-- {
		if str[i] == ' ' {
			ans = ans + str[i+1:j] + " "
			// 如果左边还是空格就继续
			for i >= 1 && str[i-1] == ' ' {
				i--
			}
			// j记录单词后的第一个空格位置
			j = i
		} else if i == l {
			ans += str[i:j]
		}
	}
	return ans
}
