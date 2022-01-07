package main

import "math/rand"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	l []int
}

func Constructor(head *ListNode) Solution {
	t := []int{}
	for head != nil {
		t = append(t, head.Val)
		head = head.Next
	}
	return Solution{
		l: t,
	}
}

func (this *Solution) GetRandom() int {
	r := rand.Intn(len(this.l))
	return this.l[r]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
