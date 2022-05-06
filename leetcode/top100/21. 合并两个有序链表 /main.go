package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := ListNode{}
	head := &res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			head.Next = list2
			break
		}

		if list2 == nil {
			head.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}

		head = head.Next
	}

	return res.Next
}
