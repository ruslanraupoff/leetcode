package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	s := 0

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, t int) {
		if root == nil {
			return
		}

		t = t*10 + root.Val

		if root.Left == nil && root.Right == nil {
			s += t
			return
		}

		dfs(root.Left, t)
		dfs(root.Right, t)
	}

	dfs(root, 0)

	return s
}
