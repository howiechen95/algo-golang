package main

/**
 *
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
func maxInWindows(num []int, size int) []int {
	// write code here
	l := len(num)

	if l == 0 || size == 0 || size > l {
		return nil
	}

	//  返回的结果集
	res := make([]int, 0)
	//  单调队列
	//  从大到小排序
	list := make([]int, 0)

	//  窗口右移，判断单调队列第一个元素是否等于被移除的元素值
	//  等于被移除的元素值，则单调队列第一个元素也移除
	pop := func(val int) {
		if list[0] == val {
			list = list[1:]
		}
	}

	//  向单调队列压入一个数据
	//  单调队列:  大于等于左边元素  小于等于右边元素
	push := func(val int) {
		for len(list) != 0 && list[len(list)-1] < val {
			list = list[:len(list)-1]
		}

		list = append(list, val)
	}

	//  前size个元素进入单调队列
	for i := 0; i < size; i++ {
		push(num[i])
	}
	//  前size个元素为第一个窗口，取第一个窗口最大值
	res = append(res, list[0])

	//  处理后续元素
	for i := size; i < len(num); i++ {
		//      处理窗口左侧被移除的元素，判断是否也需要从单调队列移除
		pop(num[i-size])
		//      添加当前窗口右侧新增的元素
		push(num[i])

		//      取单调队列最大的元素值
		res = append(res, list[0])
	}

	return res
}
