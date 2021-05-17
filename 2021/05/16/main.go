package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	ans := 0
	if dfs(root, &ans) == 2 {
		ans++
	}
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 1
	}
	l, r := dfs(root.Left, ans), dfs(root.Right, ans)
	if l == 2 || r == 2 {
		*ans++
		return 3
	} else if l == 3 || r == 3 {
		return 1
	} else {
		return 2
	}
}
