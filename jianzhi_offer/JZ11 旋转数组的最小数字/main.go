package main

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	left, right := 0, len(rotateArray)-1

	for left < right {
		middle := left + (right-left)>>1
		//如果中间值大于最右边的值，说明旋转之后最小的
		//数字肯定在mid的右边，比如[3, 4, 5, 6, 7, 1, 2]
		if rotateArray[middle] > rotateArray[right] {
			left = middle + 1
		} else if rotateArray[middle] < rotateArray[right] {
			//如果中间值小于最右边的值，说明旋转之后最小的
			//数字肯定在mid的前面，比如[6, 7, 1, 2, 3, 4, 5],
			//注意这里mid是不能减1的，比如[3，1，3]，我们这里只是
			//证明了numbers[mid]比numbers[right]小，但有可能
			//numbers[mid]是最小的，所以我们不能把它给排除掉
			right = middle
		} else {
			//如果中间值等于最后一个元素的值，我们是没法确定最小值是
			// 在mid的前面还是后面，但我们可以缩小查找范围，让right
			// 减1，因为即使right指向的是最小值，但因为他的值和mid
			// 指向的一样，我们这里并没有排除mid，所以结果是不会有影响的。
			//比如[3，1，3，3，3，3，3]和[3，3，3，3，3，1，3],中间的值
			//等于最右边的值，但我们没法确定最小值是在左边还是右边
			right--
		}
	}

	//  不断的缩小查找范围，当查找范围的长度为1的时候返回，也就是left等于right的时候
	return rotateArray[left]
}
