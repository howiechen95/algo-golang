package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param grid int整型二维数组
 * @return int整型
 */
func maxValue(grid [][]int) int {
	// write code here
	dp := new([1000][1000]int)
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x, y := i+1, j+1
			dp[x][y] = max(dp[x-1][y], dp[x][y-1]) + grid[i][j]
		}
	}
	return dp[n][m]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
