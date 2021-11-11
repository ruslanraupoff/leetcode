package main

func minStartValue(nums []int) int {
	s, mn := 0, 0
	for _, n := range nums {
		s += n
		mn = min(s, mn)
	}
	return 1 - mn
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
