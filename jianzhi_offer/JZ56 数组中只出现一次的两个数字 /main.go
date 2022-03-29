package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
func FindNumsAppearOnce(array []int) []int {
	// write code here
	l := len(array)
	if l < 2 {
		return nil
	}

	xor := 0
	for _, x := range array {
		xor ^= x
	}

	n := 1
	for i := 0; i < 64; i++ {
		if (xor & (n << i)) > 0 {
			n <<= i
			break
		}
	}

	tmp := xor
	for _, x := range array {
		if x&n > 0 {
			xor ^= x
		}
	}

	n1 := xor
	n2 := tmp ^ xor

	if n1 > n2 {
		return []int{n2, n1}
	}
	return []int{n1, n2}
}
