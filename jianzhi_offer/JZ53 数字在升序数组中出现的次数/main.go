package main

/**
 *
 * @param data int整型一维数组
 * @param k int整型
 * @return int整型
 */
func GetNumberOfK(data []int, k int) int {
	// write code here
	if len(data) == 0 {
		return 0
	}

	//  右边界没找到该数字，则左边界也不会找到，就是不存在
	//  不存在该数字，直接返回0
	right := rightBound(data, k)
	if right == -1 {
		return 0
	}

	left := leftBound(data, k)

	//  该数字右边界出现的位置 - 左边界出现的位置 + 1
	return right - left + 1
}

// k 在数组中左边界出现的位置
func leftBound(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)>>1
		if arr[middle] == k {
			right = middle - 1
		} else if arr[middle] < k {
			left = middle + 1
		} else if arr[middle] > k {
			right = middle - 1
		}
	}

	if left >= len(arr) || arr[left] != k {
		return -1
	}

	return left
}

// k 在数组中右边界出现的位置
func rightBound(arr []int, k int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := left + (right-left)>>1

		if arr[middle] == k {
			left = middle + 1
		} else if arr[middle] > k {
			right = middle - 1
		} else if arr[middle] < k {
			left = middle + 1
		}
	}

	if right < 0 || arr[right] != k {
		return -1
	}
	return right
}
