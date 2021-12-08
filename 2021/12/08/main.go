package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var tilt int
	helper(root, &tilt)
	return tilt
}

func helper(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}

	l := helper(root.Left, tilt)
	r := helper(root.Right, tilt)

	if l > r {
		*tilt += l - r
	} else {
		*tilt += r - l
	}

	return l + r + root.Val
}
