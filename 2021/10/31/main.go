package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	curr := root
	for curr != nil {
		if curr.Child == nil {
			curr = curr.Next
			continue
		}

		child := curr.Child
		for child.Next != nil {
			child = child.Next
		}
		child.Next = curr.Next
		if curr.Next != nil {
			curr.Next.Prev = child
		}
		curr.Next = curr.Child
		curr.Child.Prev = curr
		curr.Child = nil
	}
	return root
}
