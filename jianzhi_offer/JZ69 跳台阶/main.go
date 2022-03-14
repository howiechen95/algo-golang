package main

/*
//动态规划，时间On，空间On
func jumpFloor(number int ) int {
    dp := make([]int, number+1)
    dp[0], dp[1] = 1, 1

    for i := 2; i < len(dp); i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[number]
}
*/

//记忆化搜索,空间压缩为O1
func jumpFloor(number int) int {
	prev, cur := 1, 1

	for i := 2; i < number+1; i++ {
		temp := cur
		cur = prev + cur
		prev = temp
	}
	return cur
}
