package main

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 0; i < len(cost); i++ {
		c := cost[i] + Min(a, b)
		b = a
		a = c
	}
	return Min(a, b)
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
