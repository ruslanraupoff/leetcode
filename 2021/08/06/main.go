package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	d := make(map[int]int)
	level = 0
	def traverse(root, level):
		if not root:
			return
		try:
			d[level].append(root.val)
		except:
			d[level] = [root.val]            
		level += 1
		for child in root.children:
			traverse(child, level)
	traverse(root, level)
	return list(d.values())

	return r
}
