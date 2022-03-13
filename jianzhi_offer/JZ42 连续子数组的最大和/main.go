package main

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here

	return greatSumSolution2(array)
}

func greatSumSolution2(array []int) int {
	ret := array[0]
	l := len(array)
	if l == 1 {
		return ret
	}
	tmp := ret
	for i := 1; i < l; i++ {
		// 只有一个变量保存上一个的值
		tmp = max(array[i]+tmp, array[i])
		ret = max(tmp, ret)
	}
	return ret
}

func max(val1 int, val2 int) int {
	if val1 >= val2 {
		return val1
	} else {
		return val2
	}
}
