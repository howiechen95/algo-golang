package main

import "fmt"

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
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
		BubbleSort(testCase)
		fmt.Println(testCase)
	}
}
