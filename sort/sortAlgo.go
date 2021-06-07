package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		preIndex := i - 1
		current := arr[i]
		for {
			if preIndex >= 0 && arr[preIndex] > current {
				arr[preIndex+1] = arr[preIndex]
				preIndex--
			} else {
				break
			}
		}
		arr[preIndex+1] = current
	}
}

func ShellSort(arr []int) {
	length := len(arr)
	for gap := length / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < length; i++ {
			j := i
			current := arr[i]
			for {
				if j-gap >= 0 && current < arr[j-gap] {
					arr[j] = arr[j-gap]
					j = j - gap
				} else {
					break
				}
			}
			arr[j] = current
		}
	}
}

func main() {
	testCases := [][]int{
		{},
		{1},
		{1, 2},
		{2, 1},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
		{3, 4, 2, 5, 1},
	}
	for _, testCase := range testCases {
		ShellSort(testCase)
		fmt.Println(testCase)
	}
}
