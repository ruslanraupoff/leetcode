package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	prehead := &ListNode{Next: head}

	current := head
	for current != nil && current.Next != nil {
		p := current.Next
		if current.Val <= p.Val {
			current = p
			continue
		}

		current.Next = p.Next
		pre, next := prehead, prehead.Next
		for next.Val < p.Val {
			pre = next
			next = next.Next
		}
		pre.Next = p
		p.Next = next
	}

	return prehead.Next
}
