package leetcode

import "fantastic-go/algorithm/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listMap := make(map[int]*ListNode)
	dummy, cur, count := &ListNode{}, head, 0
	dummy.Next = head
	for i := 0; cur != nil; i++ {
		listMap[i] = cur
		cur = cur.Next
		count++
	}
	if count-n == 0 {
		return head.Next
	}
	cur = listMap[count-n]
	prev := listMap[count-n-1]
	prev.Next = cur.Next
	return dummy.Next
}
