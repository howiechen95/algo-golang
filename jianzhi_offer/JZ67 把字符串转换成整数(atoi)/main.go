package main

func StrToInt(str string) int {
	l := len(str)
	if l == 0 {
		return 0
	}
	sign := 1
	num := 0
	i := 0
	for ; i < l && str[i] == ' '; i++ {
	}
	if i < l && (str[i] == '+' || str[i] == '-') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	for i < l {
		if str[i] >= '0' && str[i] <= '9' {
			num = num*10 + int(str[i]-'0')
			i++
			if sign == -1 && -1*num < -2147483648 {
				return -2147483648
			} else if sign == 1 && num > 2147483647 {
				return 2147483647
			}
		} else {
			break
		}
	}
	return sign * num
}
