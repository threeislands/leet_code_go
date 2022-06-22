package src

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := ListNode{}
	current := &root
	for list1 != nil || list2 != nil {
		fmt.Print('s')
		if list2 == nil || list1.Val >= list2.Val {
			current.Next = &ListNode{Val: list1.Val}
			current = current.Next
			list1 = list1.Next
		} else {
			current.Next = &ListNode{Val: list2.Val}
			current = current.Next
			list2 = list2.Next
		}
	}
	return root.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
