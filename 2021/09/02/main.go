package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	var dfs func(int, int) []*TreeNode

	dfs = func(low, high int) []*TreeNode {
		if low > high {
			return []*TreeNode{nil}
		}
		ans := []*TreeNode{}
		for mid := low; mid <= high; mid++ {
			left := dfs(low, mid-1)
			right := dfs(mid+1, high)
			for _, l := range left {
				for _, r := range right {
					root := &TreeNode{mid, nil, nil}
					root.Left = l
					root.Right = r
					ans = append(ans, root)
				}
			}
		}
		return ans
	}
	return dfs(1, n)
}
