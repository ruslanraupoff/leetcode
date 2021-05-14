package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left == nil {
		return flatten(root.Right)
	}
	if root.Right == nil {
		root.Right = root.Left
		root.Left = nil
		return flatten(root.Right)
	}
	r := flatten(root.Right)
	flatten(root.Left).Right = root.Right
	root.Right = root.Left
	root.Left = nil
	return r
}
