package pkg

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	root := &ListNode{Next: head}
	prev, cur := root, root.Next

	for cur != nil && cur.Next != nil {
		prev.Next = cur.Next
		cur.Next = cur.Next.Next
		prev.Next.Next = cur
		prev, cur = cur, cur.Next
	}

	return root.Next
}
