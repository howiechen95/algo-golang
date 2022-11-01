package main

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for ; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] = digits[i] + 1
			return digits
		}
		digits[i] = 0
		if i == 0 {
			d := []int{1}
			d = append(d, digits...)
			return d
		}
	}
	return digits
}
