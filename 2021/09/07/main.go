package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var new *ListNode
	for head != nil {
		temp := head.Next
		head.Next = new
		new = head
		head = temp
	}
	return new
}
