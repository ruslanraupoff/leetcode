package main

import "container/heap"

func swimInWater(grid [][]int) int {
	M, N := len(grid), len(grid[0])
	buf := &IntHeap{[3]int{grid[0][0], 0, 0}}
	heap.Init(buf)
	seen := make([][]bool, M)
	for i := range seen {
		seen[i] = make([]bool, N)
	}
	height := 0

	for buf.Len() > 0 {
		c := heap.Pop(buf).([3]int)

		height = max(height, c[0])

		if c[1] == M-1 && c[2] == N-1 {
			return height
		}

		for _, da := range [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
			x, y := c[1]+da[0], c[2]+da[1]

			if x >= 0 && x < M && y >= 0 && y < N && !seen[x][y] {
				heap.Push(buf, [3]int{grid[x][y], x, y})
				seen[x][y] = true
			}
		}
	}

	return height
}

type IntHeap [][3]int

func (c IntHeap) Len() int {
	return len(c)
}

func (c IntHeap) Less(i, j int) bool {
	return c[i][0] < c[j][0]
}

func (c IntHeap) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *IntHeap) Push(x interface{}) {
	*c = append(*c, x.([3]int))
}

func (c *IntHeap) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[:n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
