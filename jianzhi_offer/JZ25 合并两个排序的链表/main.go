package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	var head = new(ListNode)

	var cur = head
	for pHead1 != nil && pHead2 != nil {
		if pHead1.Val < pHead2.Val {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		} else {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		}
		cur = cur.Next
	}
	if pHead1 == nil {
		cur.Next = pHead2
	} else {
		cur.Next = pHead1
	}
	return head.Next
}
