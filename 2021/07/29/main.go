package main

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	mn := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				continue
			}
			mat[i][j] = mn

			if 0 <= i-1 {
				mat[i][j] = min(mat[i][j], mat[i-1][j]+1)
			}
			if 0 <= j-1 {
				mat[i][j] = min(mat[i][j], mat[i][j-1]+1)
			}
		}
	}

	for i := m - 1; 0 <= i; i-- {
		for j := n - 1; 0 <= j; j-- {
			if mat[i][j] == 0 {
				continue
			}

			if i+1 < m {
				mat[i][j] = min(mat[i][j], mat[i+1][j]+1)
			}
			if j+1 < n {
				mat[i][j] = min(mat[i][j], mat[i][j+1]+1)
			}
		}
	}

	return mat
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
