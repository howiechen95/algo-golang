package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNodeš▒╗
 * @param pHead2 ListNodeš▒╗
 * @return ListNodeš▒╗
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}

	cur1, cur2 := pHead1, pHead2
	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = pHead2
		} else {
			cur1 = cur1.Next
		}

		if cur2 == nil {
			cur2 = pHead1
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}
