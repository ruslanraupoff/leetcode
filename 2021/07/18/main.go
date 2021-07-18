package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	jump := &ListNode{}
	root := jump
	r := head
	l := r
	root.Next = l

	for true {
		count := 0
		for r != nil && count < k {
			r = r.Next
			count += 1
		}
		if count == k {
			pre, cur := r, l
			for i := 0; i < k; i++ {
				cur.Next, cur, pre = pre, cur.Next, cur
			}
			jump.Next, jump, l = pre, l, r
		} else {
			return root.Next
		}
	}
	return &ListNode{}
}
