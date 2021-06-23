package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	new := &ListNode{}
	new.Next = head
	pre := new
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	curr := pre.Next
	prev := pre
	for i := 0; i <= right-left; i++ {
		prev, curr.Next, curr = curr, prev, curr.Next
	}
	pre.Next.Next = curr
	pre.Next = prev

	return new.Next
}
