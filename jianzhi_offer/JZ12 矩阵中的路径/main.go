package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param matrix char字符型二维数组
 * @param word string字符串
 * @return bool布尔型
 */
import "fmt"

//定义map来记录当前走过的路径
var m map[string]bool

func backtracking(matrix [][]byte, bytes []byte, index int, i int, j int) bool {
	if index == len(bytes) {
		return true
	} else if i >= len(matrix) || j >= len(matrix[0]) || i < 0 || j < 0 {
		m[string(i)+string(j)] = false
		return false
	} else {
		if matrix[i][j] == bytes[index] && m[string(i)+string(j)] != true {
			m[string(i)+string(j)] = true
			flag := backtracking(matrix, bytes, index+1, i+1, j) ||
				backtracking(matrix, bytes, index+1, i-1, j) ||
				backtracking(matrix, bytes, index+1, i, j+1) ||
				backtracking(matrix, bytes, index+1, i, j-1)
			if flag == false {
				m[string(i)+string(j)] = false
				return false
			}
			return true
		} else {
			m[string(i)+string(j)] = false
			fmt.Println("i=", i, "j=", j)
			return false
		}
	}
}
func hasPath(matrix [][]byte, word string) bool {
	// write code here
	m = make(map[string]bool)
	bytes := []byte(word)
	//找到矩阵中字符串首字母，从这里开始遍历
	for i, v := range matrix {
		for j, v2 := range v {
			if v2 == bytes[0] {
				if backtracking(matrix, bytes, 0, i, j) {
					return true
				}
				for i, _ := range m {
					m[i] = false
				}
			}
		}
	}
	return false
}
