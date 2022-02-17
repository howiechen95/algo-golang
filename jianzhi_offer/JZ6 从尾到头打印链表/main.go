package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListFromTailToHead(head *ListNode) []int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func printListFromTailToHead1(head *ListNode) []int {
	var arr []int
	if head != nil {
		arr = append(printListFromTailToHead1(head.Next), head.Val)
	}
	return arr
}
