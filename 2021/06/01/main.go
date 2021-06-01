package main

func maxAreaOfIsland(grid [][]int) int {
	r, n, m := 0, len(grid), len(grid[0])
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i+1, j) + dfs(i, j+1) + dfs(i-1, j) + dfs(i, j-1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				s := dfs(i, j)
				if s > r {
					r = s
				}
			}
		}
	}
	return r
}
