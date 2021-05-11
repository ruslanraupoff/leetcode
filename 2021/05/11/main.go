package main

func maxScore(cardPoints []int, k int) int {
	var t = 0
	for i := 1; i <= k; i++ {
		t += cardPoints[i-1]
	}
	var r = t
	var l = len(cardPoints)
	for i := 1; i <= k; i++ {
		t = t - cardPoints[k-i] + cardPoints[l-i]
		if t > r {
			r = t
		}
	}
	return r
}
