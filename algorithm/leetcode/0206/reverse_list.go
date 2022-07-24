package leetcode

import "CrazyCollin/personalProjects/fantastic-go/algorithm/structures"

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
