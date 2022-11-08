package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	for i := len(matrix) - 1; i >= 0; i-- {
		if matrix[i][0] <= target {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == target {
					return true
				}
			}
			return false
		}
	}
	return false
}

func main() {
	cs := []struct {
		matrix [][]int
		target int
	}{
		{
			[][]int{
				{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
			},
			3,
		},
		{
			[][]int{
				{1},
			},
			1,
		},
	}
	for _, v := range cs {
		fmt.Println(searchMatrix(v.matrix, v.target))
	}
}
