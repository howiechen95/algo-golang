package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alpM := make(map[rune]int)
	for _, v := range s {
		alpM[v]++
	}
	for _, v := range t {
		alpM[v]--
	}
	for _, v := range alpM {
		if v != 0 {
			return false
		}
	}
	return true
}
