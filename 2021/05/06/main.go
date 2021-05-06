package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return sortToBST(head, nil)
}

func sortToBST(begin *ListNode, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}
	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}
	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow
	return &TreeNode{
		Val:   mid.Val,
		Left:  sortToBST(begin, mid),
		Right: sortToBST(mid.Next, end),
	}
}
