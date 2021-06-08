package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	d := len(inorder)
	mp := make(map[int]int)
	for i := 0; i < d; i++ {
		mp[inorder[i]] = i
	}
	res, _ := toTree(preorder, mp, 0, 0, d-1)
	return res
}

func toTree(preorder []int, mp map[int]int, index, start, end int) (*TreeNode, int) {
	if start > end {
		return nil, index
	}
	cur := preorder[index]
	treeNode := &TreeNode{Val: cur}
	index += 1

	treeNode.Left, index = toTree(preorder, mp, index, start, mp[cur]-1)
	treeNode.Right, index = toTree(preorder, mp, index, mp[cur]+1, end)

	return treeNode, index
}
