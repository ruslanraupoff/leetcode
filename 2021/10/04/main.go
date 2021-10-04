package main

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func islandPerimeter(grid [][]int) int {
	r, m, n := 0, len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			r += 4
			for k := 0; k < 4; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] == 1 {
					r--
				}
			}
		}
	}

	return r
}
