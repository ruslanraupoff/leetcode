package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	i, n := 0, len(preorder)

	var bfs func(int) *TreeNode
	bfs = func(b int) *TreeNode {
		if i == n || preorder[i] > b {
			return nil
		}
		root := &TreeNode{
			Val: preorder[i],
		}
		i++
		root.Left = bfs(root.Val)
		root.Right = bfs(b)
		return root
	}

	return bfs(1 << 32)
}
