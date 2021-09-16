package main

func spiralOrder(matrix [][]int) []int {
	d := [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	r, m, n := matrix[0], len(matrix), len(matrix[0])
	i, j, p := (m-1)*n, 0, []int{0, n - 1}
	for i > 0 {
		for k := 1; k < m; k++ {
			i -= 1
			p[0] += d[j][0]
			p[1] += d[j][1]
			r = append(r, matrix[p[0]][p[1]])
		}
		m -= 1
		m, n = n, m
		j = (j + 1) % 4
	}
	return r
}
