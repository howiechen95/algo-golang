package main

func Find(target int, array [][]int) bool {
	for i := 0; i < len(array); i++ {
		for j := len(array[i]) - 1; j >= 0; j-- {
			if array[i][j] == target {
				return true
			}
		}
	}
	return false
}

func Find1(target int, array [][]int) bool {
	i := 0
	j := len(array[i]) - 1
	for i < len(array) && j >= 0 {
		if array[i][j] == target {
			return true
		} else if array[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
