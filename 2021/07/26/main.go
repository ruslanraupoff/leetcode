package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(int, int) *TreeNode
	dfs = func(l int, r int) *TreeNode {
		if l > r {
			return nil
		}

		m := l + (r-l)/2

		root := &TreeNode{Val: nums[m]}
		root.Left = dfs(l, m-1)
		root.Right = dfs(m+1, r)

		return root
	}
	return dfs(0, len(nums)-1)
}
