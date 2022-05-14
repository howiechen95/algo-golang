package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var (
		leftMax, rightMax, res int
		left, right            int
	)
	left = 0
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

func trap1(height []int) int {
	if len(height) == 0 {
		return 0
	}
	max := height[0]
	for _, i := range height {
		if i > max {
			max = i
		}
	}
	res := 0
	for i := 1; i <= max; i++ {
		tmpMin, tmpMax := -1, -1
		overCount := 0
		for j := 0; j < len(height); j++ {
			if height[j] >= i {
				overCount++
				if tmpMin == -1 {
					tmpMin = j
					tmpMax = j
				}
				if j > tmpMax {
					tmpMax = j
				}
			}
		}
		//fmt.Println("floor:", i, "overCount:", overCount, "tmpMin:", tmpMin, "tmpMax:", tmpMax, "add:", tmpMax-tmpMin-overCount+2)
		if tmpMax-tmpMin > 0 {
			res += tmpMax - tmpMin - overCount + 1
		}
	}
	return res
}

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
