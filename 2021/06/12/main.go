package main

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	d := len(stations)
	dp := make([]int, d+1)
	dp[0] = startFuel
	for i := 0; i < d; i++ {
		for j := i; j >= 0; j-- {
			if stations[i][0] <= dp[j] {
				dp[j+1] = max(dp[j+1], dp[j]+stations[i][1])
			}
		}
	}
	for i := 0; i <= d; i++ {
		if dp[i] >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
