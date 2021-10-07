package main

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(r, c, k int) bool {
		if len(word) == k {
			return true
		}

		if r < 0 || m <= r || c < 0 || n <= c || board[r][c] != word[k] {
			return false
		}

		t := board[r][c]
		board[r][c] = 0

		if dfs(r-1, c, k+1) ||
			dfs(r+1, c, k+1) ||
			dfs(r, c-1, k+1) ||
			dfs(r, c+1, k+1) {
			return true
		}

		board[r][c] = t

		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
