package main

var (
	str = ""
)

func Insert(ch byte) {
	str += string(ch)
}

func FirstAppearingOnce() byte {
	bytes := []byte(str)
	m := make(map[byte][]int)
	for i, v := range bytes {
		if m[v] == nil {
			m[v] = make([]int, 2)
		}
		m[v][0]++
		m[v][1] = i
	}
	index := len(bytes)
	for _, v := range m {
		if v[0] == 1 && index > v[1] {
			index = v[1]
		}
	}
	if index == len(bytes) {
		return '#'
	} else {
		return bytes[index]
	}

}
