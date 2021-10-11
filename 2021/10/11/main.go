package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	_, d := helper(root)
	return d
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	ll, ld := helper(root.Left)
	rl, rd := helper(root.Right)

	return max(ll, rl) + 1, max(ll+rl, max(ld, rd))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
