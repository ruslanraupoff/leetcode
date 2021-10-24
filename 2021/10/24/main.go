package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	r := 0
	dfs(root, &r)
	return r
}

func dfs(root *TreeNode, r *int) {
	if root == nil {
		return
	}
	*r++
	dfs(root.Left, r)
	dfs(root.Right, r)
}
