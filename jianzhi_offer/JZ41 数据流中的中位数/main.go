package main

import (
	"fmt"
	"sort"
)

var arr = make([]int, 0)

func Insert(num int) {
	arr = append(arr, num)
}

func GetMedian() float64 {
	length := len(arr)
	if length == 1 {
		return float64(arr[0])
	}
	sort.Ints(arr)
	midIndex := length / 2
	if length%2 == 0 {
		return float64(arr[midIndex-1]+arr[midIndex]) / 2
	} else {
		return float64(arr[midIndex])
	}
}

func main() {
	Insert(1)
	fmt.Println(GetMedian())
	Insert(2)
	fmt.Println(GetMedian())
	Insert(4)
	fmt.Println(GetMedian())
}
