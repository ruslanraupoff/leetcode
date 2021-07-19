package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
