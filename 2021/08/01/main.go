package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dfs(grid [][]int, r int, c int, visited map[[2]int]bool, clr int, count *int) {

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	visited[[2]int{r, c}] = true

	grid[r][c] = clr
	*count = *count + 1

	for _, dir := range dirs {
		nr := r + dir[0]
		nc := c + dir[1]

		if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[0]) && !visited[[2]int{nr, nc}] && grid[nr][nc] == 1 {
			dfs(grid, nr, nc, visited, clr, count)
		}
	}
}

func largestIsland(grid [][]int) int {

	islandSize := make(map[int]int)
	zeroPresent := false
	paint := 2
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				visited := make(map[[2]int]bool)
				count := 0
				dfs(grid, r, c, visited, paint, &count)
				islandSize[paint] = count
				paint++
			} else if grid[r][c] == 0 {
				zeroPresent = true
			}
		}
	}

	if !zeroPresent {
		return len(grid) * len(grid[0])
	}

	maxSize := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 0 {
				size := 0
				checked := make(map[int]bool)

				dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

				for _, dir := range dirs {
					nr := r + dir[0]
					nc := c + dir[1]

					if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[0]) {
						paint := grid[nr][nc]
						if !checked[paint] {
							size += islandSize[paint]
							checked[paint] = true
						}
					}
					maxSize = max(maxSize, size+1)
				}
			}
		}
	}

	return maxSize
}
