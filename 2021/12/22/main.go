package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	cur := head
	size := 0
	for cur != nil {
		cur = cur.Next
		size++
	}

	cur = head
	for i := 0; i < (size-1)/2; i++ {
		cur = cur.Next
	}

	next := cur.Next
	for next != nil {
		temp := next.Next
		next.Next = cur
		cur = next
		next = temp
	}
	end := cur

	for head != end {
		hnext := head.Next
		enext := end.Next
		head.Next = end
		end.Next = hnext
		head = hnext
		end = enext
	}

	end.Next = nil
}
