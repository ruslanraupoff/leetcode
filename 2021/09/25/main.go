package main

type state struct {
	x int
	y int
	z int
}

func shortestPath(grid [][]int, k int) int {
	s, n, m := 0, len(grid), len(grid[0])
	l := [][]int{}
	l = append(l, []int{0, 0, k})
	v := map[state]bool{}
	v[state{0, 0, k}] = true
	rows := []int{-1, 0, 1, 0}
	cols := []int{0, -1, 0, 1}

	for len(l) > 0 {
		ns := [][]int{}
		for _, xyr := range l {
			x, y, r := xyr[0], xyr[1], xyr[2]

			if x == n-1 && y == m-1 {
				return s
			}

			for t := range rows {
				i, j := x+rows[t], y+cols[t]
				if i >= 0 && j >= 0 && i < n && j < m {
					if grid[i][j] == 0 {
						d := state{i, j, r}
						if !isVisited(d, v) {
							v[d] = true
							ns = append(ns, []int{i, j, r})
						}
					} else if r > 0 {
						d := state{i, j, r - 1}
						if !isVisited(d, v) {
							v[d] = true
							ns = append(ns, []int{i, j, r - 1})
						}
					}
				}
			}
		}

		l = ns
		s += 1
	}

	return -1
}

func isVisited(d state, visited map[state]bool) bool {
	if _, exists := visited[d]; exists {
		return true
	}
	return false
}
