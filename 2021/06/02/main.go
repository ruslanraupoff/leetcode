package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2 := len(s1), len(s2)
	if l1+l2 != len(s3) {
		return false
	}
	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
		for j := 0; j <= l2; j++ {
			dp[i][j] = false
		}
	}
	dp[0][0] = true
	for i := 1; i <= l1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}
	return dp[l1][l2]
}
