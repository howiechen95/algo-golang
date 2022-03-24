package main

var ans int = 0
var mod int = 1000000007

func merge(data []int, l, mid, r int) {
	tmp := make([]int, r-l+1)
	idx := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r {
		if data[i] <= data[j] {
			tmp[idx] = data[i]
			i++
			idx++
		} else {
			ans += (mid - i + 1)
			ans %= mod
			tmp[idx] = data[j]
			j++
			idx++
		}
	}
	for i <= mid {
		tmp[idx] = data[i]
		idx++
		i++
	}
	for j <= r {
		tmp[idx] = data[j]
		idx++
		j++
	}
	for k := 0; k < r-l+1; k++ {
		data[l+k] = tmp[k]
	}
}
func mergeSort(data []int, l, r int) {
	if l < r {
		mid := (l + r) >> 1
		mergeSort(data, l, mid)
		mergeSort(data, mid+1, r)
		merge(data, l, mid, r)
	}
}
func InversePairs(data []int) int {
	// write code here
	mergeSort(data, 0, len(data)-1)
	return ans
}
