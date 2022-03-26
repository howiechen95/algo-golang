package main

//修改快排，平均nlogn,最坏On^2 ; 空间平均logn，最坏On
func GetLeastNumbers_Solution(input []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	quickSort(input, 0, len(input)-1)
	return input[:k]
}

func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}

	i, j, base := left, right, nums[left] //优化点，随机选基准

	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
