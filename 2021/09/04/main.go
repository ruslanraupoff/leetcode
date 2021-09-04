package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	tree := make([][]int, n)
	for _, e := range edges {
		i, j := e[0], e[1]
		tree[i] = append(tree[i], j)
		tree[j] = append(tree[j], i)
	}

	res := make([]int, n)
	count := make([]int, n)

	isSeen := make([]bool, n)
	deepFirstSearch(0, isSeen, count, res, tree)

	isSeen = make([]bool, n)
	deepFirstMove(0, isSeen, count, res, tree)

	return res
}

func deepFirstSearch(root int, isSeen []bool, count, res []int, tree [][]int) {
	isSeen[root] = true
	for _, n := range tree[root] {
		if isSeen[n] {
			continue
		}
		deepFirstSearch(n, isSeen, count, res, tree)
		count[root] += count[n]
		res[root] += res[n] + count[n]
	}
	count[root]++
}

func deepFirstMove(root int, isSeen []bool, count, res []int, tree [][]int) {
	N := len(isSeen)
	isSeen[root] = true
	for _, n := range tree[root] {
		if isSeen[n] {
			continue
		}
		res[n] = res[root] + (N - count[n]) - count[n]
		deepFirstMove(n, isSeen, count, res, tree)
	}
}
