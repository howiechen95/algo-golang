package main

import "sort"

func Search(n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}
func search(nums []int, target int) int {
	// 二分规则：当 nums[i] 在 target 左侧时返回 false，否则返回 true
	i := sort.Search(len(nums), func(i int) bool {
		if nums[i] >= nums[0] { // nums[i] 落在左半部分
			return target >= nums[0] && nums[i] >= target // 如果 target 在右半部分可以直接返回 false
		} else { // nums[i] 落在右半部分
			return target >= nums[0] || nums[i] >= target // 如果 target 在左半部分可以直接返回 true
		}
	})
	if i < len(nums) && nums[i] == target {
		return i
	}
	return -1
}
