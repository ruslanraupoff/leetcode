package main

import "sort"

func makesquare(matchsticks []int) bool {
	per := 0
	for _, value := range matchsticks {
		per += value
	}
	if per%4 != 0 {
		return false
	}
	side := per / 4
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	sq := [4]int{}

	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return sq[0] == sq[1] && sq[1] == sq[2] && sq[2] == side
		}

		for i := 0; i < 4; i++ {
			if sq[i]+matchsticks[idx] <= side {
				sq[i] += matchsticks[idx]
				if dfs(idx + 1) {
					return true
				}
				sq[i] -= matchsticks[idx]
			}
		}

		return false
	}

	return dfs(0)
}
