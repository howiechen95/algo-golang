package main

/**
 *
 * @param index int整型
 * @return int整型
 */
func GetUglyNumber_Solution(index int) int {
	if index <= 6 {
		return index
	}
	var i2, i3, i5 int
	res := make([]int, index)
	res[0] = 1
	for i := 1; i < index; i++ {
		res[i] = min(res[i2]*2, min(res[i3]*3, res[i5]*5))
		if res[i] == res[i2]*2 {
			i2++
		}
		if res[i] == res[i3]*3 {
			i3++
		}
		if res[i] == res[i5]*5 {
			i5++
		}
	}
	return res[index-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
