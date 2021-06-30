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
	var findLCA func(*TreeNode, *TreeNode, *TreeNode) *TreeNode

	findLCA = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}

		l := findLCA(root.Left, p, q)
		r := findLCA(root.Right, p, q)

		if l != nil && r != nil {
			return root
		}
		if l != nil {
			return l
		}
		return r
	}
	return findLCA(root, p, q)
}
