package problem0237

import "github.com/mogu-echo/go-algorithm/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
