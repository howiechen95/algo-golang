package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var pNewHead *ListNode
	for pHead != nil {
		pNext := pHead.Next
		pHead.Next = pNewHead
		pNewHead = pHead
		pHead = pNext
	}
	return pNewHead
}
