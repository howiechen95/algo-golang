package main

func moveZeroes(nums []int) {
	zeroCnt := 0
	curIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCnt++
		} else {
			nums[curIdx] = nums[i]
			curIdx++
		}
	}
	for j := len(nums) - zeroCnt; j <= len(nums)-1; j++ {
		nums[j] = 0
	}
}
