package main

import "math"

func maxProfit(prices []int) int {
	l, r := math.MaxInt32, math.MaxInt32
	p, t := 0, 0
	for _, x := range prices {
		l = min(x, l)
		p = max(x-l, p)
		r = min(x-p, r)
		t = max(x-r, t)
	}
	return t
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
