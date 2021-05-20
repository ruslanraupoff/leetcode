package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var l [][]int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, i int) {
		if root == nil {
			return
		}

		if len(l) < i+1 {
			l = append(l, []int{})
		}

		l[i] = append(l[i], root.Val)

		dfs(root.Left, i+1)
		dfs(root.Right, i+1)
	}

	dfs(root, 0)

	return l
}
