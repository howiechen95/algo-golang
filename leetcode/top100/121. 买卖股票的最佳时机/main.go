package main

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt
	res := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > res {
			res = prices[i] - min
		}
	}
	return res
}
