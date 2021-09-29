package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := length(head)
	md := n % k
	res := make([]*ListNode, k)
	i := 0
	for {
		res[i] = head
		i++
		if i == k {
			break
		}

		l := n / k
		if md > 0 {
			l++
			md--
		}

		for l > 1 && head != nil {
			head = head.Next
			l--
		}

		if head == nil {
			break
		}

		head.Next, head = nil, head.Next
	}

	return res
}

func length(head *ListNode) int {
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
