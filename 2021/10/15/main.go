package main

func maxProfit(prices []int) int {
	n := len(prices)

	b := make([]int, n+1)
	b[1] = 0 - prices[0]
	s := make([]int, n+1)

	for i := 2; i <= n; i++ {
		b[i] = max(b[i-1], s[i-2]-prices[i-1])
		s[i] = max(s[i-1], b[i-1]+prices[i-1])
	}

	return s[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
