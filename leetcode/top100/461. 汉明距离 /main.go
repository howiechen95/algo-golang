package main

func hammingDistance(x int, y int) int {
	res := x ^ y
	cnt := 0
	for res != 0 {
		if res&1 == 1 {
			cnt++
		}
		res = res >> 1
	}
	return cnt
}
