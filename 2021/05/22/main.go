package main

func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{}
	}

	board := newBoard(n + 1)
	ans := make([][]string, 0)

	backtrack(board, 1, n, &ans)

	return ans
}

func backtrack(board [][]int, row, n int, res *[][]string) {
	if row == n+1 {
		*res = append(*res, outputResult(board, n))
		return
	}

	for i := 1; i <= n; i++ {
		if canOccupy(board, row, i, n) {
			board[row][i] = 1
			backtrack(board, row+1, n, res)
			board[row][i] = 0
		}
	}
}

func canOccupy(board [][]int, row, col, n int) bool {
	for j := 1; j < row; j++ {
		for k := 1; k <= n; k++ {
			if board[j][k] == 1 {
				if row == j {
					return false
				}
				if col == k {
					return false
				}
				if abs(row-j) == abs(col-k) {
					return false
				}
			}
		}
	}
	return true
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func outputResult(board [][]int, n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		var tmp string
		for j := 1; j <= n; j++ {
			if board[i][j] == 0 {
				tmp += "."
			} else {
				tmp += "Q"
			}
		}
		res = append(res, tmp)
	}
	return res
}

func newBoard(n int) [][]int {
	var board [][]int
	for i := 0; i < n; i++ {
		board = append(board, make([]int, n))
	}
	return board
}
