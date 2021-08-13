package main

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	f := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			f = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}

	for i := m - 1; i > -1; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if f {
			matrix[i][0] = 0
		}
	}
}
