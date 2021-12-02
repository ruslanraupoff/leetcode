package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	oe := head
	eb := head.Next
	ee := eb

	for ee != nil && ee.Next != nil {
		oe.Next = ee.Next
		oe = oe.Next
		ee.Next = oe.Next
		ee = ee.Next
		oe.Next = eb
	}

	return head
}
