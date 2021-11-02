package main

var dx = []int{0, 0, 1, -1}
var dy = []int{1, -1, 0, 0}

func uniquePathsIII(grid [][]int) int {

	m, n := len(grid), len(grid[0])

	queue := make([][2]int, 0, m*n*10)
	paths := make([]int, 0, m*n*10)
	walked := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 0:
				walked |= 1 << uint(i*n+j)
			case 1:
				walked |= 1 << uint(i*n+j)
				queue = append(queue, [2]int{i, j})
				paths = append(paths, 1<<uint(i*n+j))
			}
		}
	}

	count := 0

	for len(queue) > 0 {
		size := len(queue)
		for s := 0; s < size; s++ {
			x, y := queue[s][0], queue[s][1]
			path := paths[s]
			for k := 0; k < 4; k++ {
				i, j := x+dx[k], y+dy[k]
				if i < 0 || m <= i ||
					j < 0 || n <= j ||
					path&(1<<uint(i*n+j)) != 0 {
					continue
				}
				switch grid[i][j] {
				case 0:
					queue = append(queue, [2]int{i, j})
					paths = append(paths, path|(1<<uint(i*n+j)))
				case 2:
					if path == walked {
						count++
					}
				}
			}
		}
		queue = queue[size:]
		paths = paths[size:]
	}

	return count
}
