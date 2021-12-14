package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	s := 0
	if root == nil {
		return 0
	}

	if root.Val < low {
		s = rangeSumBST(root.Right, low, high)
	} else if root.Val > high {
		s = rangeSumBST(root.Left, low, high)
	} else {
		s = root.Val
		s += rangeSumBST(root.Left, low, high)
		s += rangeSumBST(root.Right, low, high)
	}

	return s
}
