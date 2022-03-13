package main

func FindGreatestSumOfSubArray(array []int) []int {
	dp := make([]int, len(array))
	dp[0] = array[0]
	var left, right, snapLeft, snapRight int
	maxLength, maxSum := 1, array[0]
	for i := 1; i <= len(array)-1; i++ {
		right++
		dp[i] = Max(array[i]+dp[i-1], array[i])
		if array[i]+dp[i-1] < array[i] {
			left = right
		}
		if dp[i] > maxSum || dp[i] == maxSum && (right-left+1) > maxLength {
			snapLeft = left
			snapRight = right
			maxLength = right - left + 1
			maxSum = dp[i]
		}
	}
	res := make([]int, maxLength)
	idx := 0
	for i := snapLeft; i <= snapRight; i++ {
		res[idx] = array[i]
		idx++
	}
	return res
}

//Max max
func Max(a, b int) int {
	if a > b {
		return a
	}
	if b > a {
		return b
	}
	return a
}
