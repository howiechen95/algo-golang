package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型一维数组
 * @return int整型一维数组
 */
import "fmt"

func reOrderArray(array []int) []int {
	// write code here
	res := make([]int, len(array))
	head, ptr_head := 0, 0
	tail, ptr_tail := len(array)-1, len(array)-1
	for ptr_head <= len(array)-1 && ptr_tail >= 0 {
		if array[ptr_head]%2 == 1 {
			res[head] = array[ptr_head]
			head++
		}
		if array[ptr_tail]%2 == 0 {
			res[tail] = array[ptr_tail]
			tail--
		}
		ptr_head++
		ptr_tail--
	}
	fmt.Printf("%v", res)
	return res
}
