package main

func myAtoi(s string) int {
	min32Int, max32Int := -1<<31, 1<<31-1
	i, flag, sign, result := 0, false, 1, 0
	l := len(s)
	for ; i < l && s[i] == ' '; i++ {
	}
	if l == i {
		return 0
	}

	for ; i < l; i++ {
		k := string(rune(s[i]))
		if k == "-" && flag == false {
			sign = -1
			flag = true
		} else if k == "+" && flag == false {
			sign = 1
			flag = true
		} else if k == "." {
			return result
		} else if k >= "0" && k <= "9" {
			v := int(s[i] - '0')
			if flag == false {
				if v < 0 {
					sign = -1
				}
				flag = true
			}
			result = sign * (sign*result*10 + v)
		} else {
			break
		}

		if sign == -1 && result < min32Int {
			return min32Int
		}
		if sign == 1 && result > max32Int {
			return max32Int
		}
	}

	return result
}
