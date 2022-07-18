package main

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		if len(dp) == 0 || dp[len(dp)-1] < num {
			dp = append(dp, num)
		} else {
			l, r := 0, len(dp)-1
			pos := r
			for l <= r {
				mid := (l + r) >> 1
				if dp[mid] >= num {
					pos = mid
					r = mid - 1
				} else {
					l = mid + 1
				}
			}
			dp[pos] = num
		} //二分查找
	}
	return len(dp)
}
