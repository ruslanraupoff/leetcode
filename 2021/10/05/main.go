package main

func climbStairs(n int) int {
	a, b, c := 1, 1, 1
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
