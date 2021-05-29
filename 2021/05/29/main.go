package main

func totalNQueens(n int) int {
	s := 0
	dfs(&s, n, 0, 0, 0, 0)
	return s
}

func dfs(s *int, n, r, c, p, a int) {
	if r == n {
		(*s)++
		return
	}

	bits := (^(c | p | a)) & (1<<n - 1)

	for bits != 0 {
		pos := bits & (-bits)
		bits = bits & (bits - 1)
		dfs(s, n, r+1, c|pos, (p|pos)<<1, (a|pos)>>1)
	}
}
