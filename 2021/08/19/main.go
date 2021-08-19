package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxProduct(root *TreeNode) int {
	r := 0
	m := 1000000007
	s := []int{}
	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		s = append(s, root.Val+dfs(root.Left)+dfs(root.Right))
		return s[len(s)-1]
	}

	t := dfs(root)

	for _, x := range s {
		r = max(r, x*(t-x))
	}
	return r % m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
