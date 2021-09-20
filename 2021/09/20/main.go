package main

func tictactoe(moves [][]int) string {
	grid := [3][3]int{}
	for i, move := range moves {
		grid[move[0]][move[1]] = i&1 + 1
	}

	for i := 0; i < 3; i++ {
		if player := grid[i][i]; player != 0 {
			if (grid[i][0] == grid[i][1] && grid[i][0] == grid[i][2]) ||
				(grid[0][i] == grid[1][i] && grid[0][i] == grid[2][i]) {
				return string(byte('@' + player))
			}
		}
	}

	if player := grid[1][1]; player != 0 {
		if (grid[0][0] == grid[2][2] && grid[0][0] == player) ||
			(grid[2][0] == grid[0][2] && grid[2][0] == player) {
			return string(byte('@' + player))
		}
	}

	if len(moves) != 9 {
		return "Pending"
	}

	return "Draw"
}
