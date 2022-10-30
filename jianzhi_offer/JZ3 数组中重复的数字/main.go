package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func duplicate(numbers []int) int {
	count := make([]int, 10001)
	for _, n := range numbers {
		if count[n] > 0 {
			return n
		}
		count[n]++
	}
	return -1
}

func findRepeatNumber(nums []int) int {
	var a = nums[0]
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			a = nums[i]
			break
		}
	}

	for {
		fmt.Println(nums)
		if nums[a] == a {
			return a
		}
		a, nums[a] = nums[a], a
	}
	return a
}

func main() {
	testCases := [][]int{
		//{0, 1, 2, 3, 4, 5, 5},
		//{5, 4, 3, 2, 1, 1, 0},
		//{3, 0, 4, 4, 2},
		{3, 4, 2, 0, 0, 1},
	}
	for _, v := range testCases {
		fmt.Println(findRepeatNumber(v))
	}
}
