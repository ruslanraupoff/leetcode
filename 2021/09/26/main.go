package main

func movesToChessboard(board [][]int) int {
	n := len(board)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}

	rsm, csm, rsp, csp := 0, 0, 0, 0

	for i := 0; i < n; i++ {
		rsm += board[0][i]
		csm += board[i][0]
		if board[i][0] == i%2 {
			rsp++
		}
		if board[0][i] == i%2 {
			csp++
		}
	}

	if rsm < n/2 || (n+1)/2 < rsm {
		return -1
	}

	if csm < n/2 || (n+1)/2 < csm {
		return -1
	}

	if n%2 == 1 {
		if csp%2 == 1 {
			csp = n - csp
		}
		if rsp%2 == 1 {
			rsp = n - rsp
		}
	} else {
		csp = min(n-csp, csp)
		rsp = min(n-rsp, rsp)
	}
	return (csp + rsp) / 2
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
