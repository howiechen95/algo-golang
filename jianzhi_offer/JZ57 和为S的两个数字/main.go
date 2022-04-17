package main

import (
	"fmt"
	"math"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param array int整型一维数组
 * @param sum int整型
 * @return int整型一维数组
 */
func FindNumbersWithSum(array []int, sum int) []int {
	if len(array) == 0 {
		return []int{}
	}
	ret := make([]int, 2)
	tmp := math.MaxInt
	i, j := 0, len(array)
	ijFlag := true
	for i < j {
		if array[i]+array[j-1] == sum {
			if array[i]*array[j-1] < tmp {
				tmp = array[i] * array[j-1]
				ret[0], ret[1] = array[i], array[j-1]
				ijFlag = i == j-1
			}
			i++
			j--
		} else if array[i]+array[j-1] < sum {
			i++
		} else {
			j--
		}
	}
	if ret[0] == ret[1] && ijFlag {
		return []int{}
	} else {
		return ret
	}
}

func main() {
	fmt.Println(FindNumbersWithSum([]int{1, 2, 4, 7, 11, 16}, 10))

}
