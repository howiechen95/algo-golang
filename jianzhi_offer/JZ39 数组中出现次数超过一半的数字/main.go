package main

/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	t := make(map[int]int)
	for _, v := range numbers {
		t[v]++
	}
	for v, _ := range t {
		if t[v] > (len(numbers) / 2) {
			return v
		}
	}
	return -1
}
