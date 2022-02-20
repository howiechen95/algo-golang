package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	// write code here
	if pHead == nil {
		return nil
	}
	fast := pHead
	for i := 0; i < k; i++ {
		if fast == nil { //注意有可能给的k大于链表长度
			return fast
		}
		if fast != nil {
			fast = fast.Next
		}
	}
	last := pHead
	for fast != nil {
		fast = fast.Next
		last = last.Next
	}
	return last
}
