package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	helper(root)
	return root
}

func helper(node *Node) {
	if node == nil {
		return
	}
	if node.Left != nil {
		node.Left.Next = node.Right
	}
	if node.Next != nil && node.Left != nil {
		node.Right.Next = node.Next.Left
	}
	helper(node.Left)
	helper(node.Right)
}
