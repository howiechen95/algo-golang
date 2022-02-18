package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//双指针法
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return nil
	}

	fast, slow := pHead, pHead
	for fast != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			tmp := pHead
			for tmp != slow {
				tmp = tmp.Next
				slow = slow.Next
			}
			return tmp
		}
	}
	return nil
}
