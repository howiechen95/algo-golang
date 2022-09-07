package main

import "fmt"

func searchInsert(nums []int, target int) int {
	lIdx, rIdx := 0, len(nums)-1
	for lIdx < rIdx {
		mIdx := (lIdx + rIdx) / 2
		if nums[mIdx] == target {
			return mIdx
		}
		if nums[mIdx] < target {
			lIdx = mIdx + 1
		} else {
			rIdx = mIdx
		}
	}
	if nums[lIdx] < target {
		lIdx++
	}
	return lIdx
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		res    int
	}{
		//{
		//	[]int{1, 3, 5, 6},
		//	2,
		//	1,
		//},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
	}

	for _, v := range testCases {
		fmt.Println(searchInsert(v.nums, v.target) == v.res)
	}
}
