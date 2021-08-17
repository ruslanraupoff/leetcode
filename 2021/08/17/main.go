package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	var dfs func(*TreeNode, int) int

	dfs = func(root *TreeNode, mx int) int {
		if root == nil {
			return 0
		}

		cur := 0
		if root.Val >= mx {
			cur = 1
			mx = root.Val
		}

		return dfs(root.Left, mx) + dfs(root.Right, mx) + cur
	}
	return dfs(root, 10000001)
}
