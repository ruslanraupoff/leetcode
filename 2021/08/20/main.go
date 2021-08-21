package main

func isValidSudoku(board [][]byte) bool {
	for r := 0; r < 9; r++ {
		if !isValidSudokuRow(board, r) {
			return false
		}
	}

	for c := 0; c < 9; c++ {
		if !isValidSudokuCol(board, c) {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if !isValidSudokuPod(board, i) {
			return false
		}
	}
	return true
}

func isValidSudokuRow(board [][]byte, row int) bool {
	var nums [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuCol(board [][]byte, col int) bool {
	var nums [10]bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if nums[n] {
			return false
		}
		nums[n] = true
	}
	return true
}

func isValidSudokuPod(board [][]byte, pod int) bool {
	var nums [10]bool

	row := (pod / 3) * 3
	col := (pod % 3) * 3

	for drow := 0; drow < 3; drow++ {
		for dcol := 0; dcol < 3; dcol++ {
			n := convertToNumber(board[row+drow][col+dcol])
			if n < 0 {
				continue
			}
			if nums[n] {
				return false
			}
			nums[n] = true
		}
	}
	return true
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}
