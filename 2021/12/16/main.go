package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	ns := make([][]int, n)
	for i := 0; i < n; i++ {
		ns[i] = make([]int, 0, 5)
	}

	cn := make([]int, n)
	for _, e := range edges {
		ns[e[0]] = append(ns[e[0]], e[1])
		ns[e[1]] = append(ns[e[1]], e[0])
		cn[e[0]]++
		cn[e[1]]++
	}

	var ls = make([]int, 0, n)
	for i := 0; i < n; i++ {
		if cn[i] == 1 {
			ls = append(ls, i)
		}
	}

	var nd, leaf int

	for n > 2 {
		nls := make([]int, 0, len(ls))
		for _, leaf = range ls {
			n--
			for _, nd = range ns[leaf] {
				cn[nd]--
				if cn[nd] == 1 {
					nls = append(nls, nd)
				}
			}
		}
		ls = nls
	}

	var res = make([]int, 0, len(ls))
	for _, v := range ls {
		res = append(res, v)
	}

	return res
}
