package main

func digitalSum(num int) int {
	var sum int
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func dfs(threshold int, rows int, cols int, x int, y int, isVisited [][]bool) int {
	if x < 0 || y < 0 || x >= rows || y >= cols || isVisited[x][y] || digitalSum(x)+digitalSum(y) > threshold {
		return 0
	}
	isVisited[x][y] = true
	return 1 + dfs(threshold, rows, cols, x-1, y, isVisited) + dfs(threshold, rows, cols, x+1, y, isVisited) + dfs(threshold, rows, cols, x, y-1, isVisited) + dfs(threshold, rows, cols, x, y+1, isVisited)
}

func movingCount(threshold int, rows int, cols int) int {
	if threshold < 0 || rows <= 0 || cols <= 0 {
		return 0
	}
	isVisited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		isVisited[i] = make([]bool, cols)
	}
	return dfs(threshold, rows, cols, 0, 0, isVisited)
}
