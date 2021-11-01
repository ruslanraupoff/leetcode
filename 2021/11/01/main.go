package main

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}

	n := len(board[0])
	if n <= 2 {
		return
	}

	isEscaped := make([][]bool, m)
	for i := 0; i < m; i++ {
		isEscaped[i] = make([]bool, n)
	}

	idxM := make([]int, 0, m*n)
	idxN := make([]int, 0, m*n)

	bfs := func(i, j int) {
		isEscaped[i][j] = true
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			i := idxM[0]
			j := idxN[0]
			idxM = idxM[1:]
			idxN = idxN[1:]

			if 0 <= i-1 && board[i-1][j] == 'O' && !isEscaped[i-1][j] {
				idxM = append(idxM, i-1)
				idxN = append(idxN, j)
				isEscaped[i-1][j] = true
			}

			if 0 <= j-1 && board[i][j-1] == 'O' && !isEscaped[i][j-1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j-1)
				isEscaped[i][j-1] = true
			}

			if i+1 < m && board[i+1][j] == 'O' && !isEscaped[i+1][j] {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
				isEscaped[i+1][j] = true
			}

			if j+1 < n && board[i][j+1] == 'O' && !isEscaped[i][j+1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
				isEscaped[i][j+1] = true
			}
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !isEscaped[0][j] {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' && !isEscaped[m-1][j] {
			bfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !isEscaped[i][0] {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' && !isEscaped[i][n-1] {
			bfs(i, n-1)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !isEscaped[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
