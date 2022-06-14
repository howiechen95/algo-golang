package main

import "fmt"

func singleNumber(nums []int) int {
	r := nums[0]
	for i := 1; i < len(nums); i++ {
		r = nums[i] ^ r
	}
	return r
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 3, 1, 2, 3, 4}))
}
