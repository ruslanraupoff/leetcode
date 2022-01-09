package main

func isRobotBounded(instructions string) bool {
	step := [][]int{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}
	c, x, y := 0, 0, 0
	for _, i := range instructions {
		if i == 'G' {
			x += step[c][0]
			y += step[c][1]
		} else if i == 'L' {
			c = (c - 1) % 4
		} else if i == 'R' {
			c = (c + 1) % 4
		}
	}
	return (x == 0 && y == 0) || c != 0
}
