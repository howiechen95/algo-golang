package main

func findDuplicate(nums []int) int {
	n := len(nums)
	ans := 0
	bit_max := 31
	for ((n - 1) >> bit_max) == 0 {
		bit_max--
	}
	for bit := 0; bit <= bit_max; bit++ {
		x, y := 0, 0
		for i := 0; i < n; i++ {
			if (nums[i] & (1 << bit)) > 0 {
				x++
			}
			if i >= 1 && (i&(1<<bit)) > 0 {
				y++
			}
		}
		if x > y {
			ans |= 1 << bit
		}
	}
	return ans
}
