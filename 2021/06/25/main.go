package main

func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	p := make([]int, n+1)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	var ifind func(int) int
	ifind = func(i int) int {
		for i != p[i] {
			i = p[i]
		}
		return i
	}

	var i int
	var e []int
	for i, e = range edges {
		f, t := e[0], e[1]
		pf := ifind(f)
		pt := ifind(t)
		if pf == pt {
			break
		}
		p[pf] = pt
	}
	return edges[i]
}
