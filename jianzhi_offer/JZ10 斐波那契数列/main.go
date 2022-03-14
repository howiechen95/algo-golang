package main

/**
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}

	temp := 0
	a, b := 1, 1

	for i := 2; i < n; i++ {
		temp = a + b
		a = b
		b = temp
	}
	return temp
}
