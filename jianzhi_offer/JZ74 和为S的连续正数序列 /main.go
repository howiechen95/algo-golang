package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param sum int整型
 * @return int整型二维数组
 */
func FindContinuousSequence(sum int) [][]int {
	res := make([][]int, 0)
	tmp := 0
	for i := 1; i <= sum/2; i++ {
		for j := i; j < sum; j++ {
			tmp += j
			if tmp == sum {
				ans := make([]int, 0)
				for k := i; k <= j; k++ {
					ans = append(ans, k)
				}
				res = append(res, ans)
			} else if tmp > sum {
				tmp = 0
				break
			}
		}
	}
	return res
}
