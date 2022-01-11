package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, b int) int {
	if node == nil {
		return 0
	}
	b = b*2 + node.Val
	if node.Left == node.Right {
		return b
	}
	return dfs(node.Left, b) + dfs(node.Right, b)
}
