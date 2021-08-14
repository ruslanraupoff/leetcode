package main

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	return remove(boxes, dp, 0, n-1, 0)
}

func remove(boxes []int, dp [][][]int, l, r, k int) int {
	if l > r {
		return 0
	}

	if dp[l][r][k] > 0 {
		return dp[l][r][k]
	}

	for l < r && boxes[r-1] == boxes[r] {
		r--
		k++
	}

	res := (k+1)*(k+1) + remove(boxes, dp, l, r-1, 0)

	for m := r - 1; l <= m; m-- {
		if boxes[r] == boxes[m] {
			res = max(res, remove(boxes, dp, m+1, r-1, 0)+remove(boxes, dp, l, m, k+1))
		}
	}

	dp[l][r][k] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
