package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	return helper(root, root, k)
}
func helper(root, sroot *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	return (2*root.Val != k && bs(sroot, k-root.Val)) || helper(root.Left, sroot, k) || helper(root.Right, sroot, k)
}
func bs(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if root.Val == k {
		return true
	}
	if root.Val < k {
		return bs(root.Right, k)
	}
	return bs(root.Left, k)
}
