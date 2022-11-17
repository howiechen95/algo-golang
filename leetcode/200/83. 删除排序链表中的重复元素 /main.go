package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := head
	cur, next := head, head.Next
	for next != nil {
		if cur.Val != next.Val {
			cur.Next = next
			cur = next
			next = next.Next
		} else {
			cur.Next = next.Next
			next = next.Next
		}
	}
	return res
}
