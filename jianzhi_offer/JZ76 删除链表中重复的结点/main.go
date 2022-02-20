package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplication(pHead *ListNode) *ListNode {
	var minVal int = -999
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	head := &ListNode{
		Val: minVal,
	}
	head.Next = pHead
	pre, cur := head, head.Next
	for cur != nil {
		if cur.Next != nil && cur.Next.Val == cur.Val {
			// 相同结点一直前进
			for cur.Next != nil && cur.Next.Val == cur.Val {
				cur = cur.Next
			}
			// 退出循环时，cur 指向重复值，也需要删除，而 cur.next 指向第一个不重复的值
			// cur 继续前进
			cur = cur.Next
			// pre 连接新结点
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head.Next
}
