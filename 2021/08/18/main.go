package main

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp = append(dp, 0)
	}

	dp[0], dp[1] = 1, 1
	l, r := int(s[0]-'0'), 0
	for i := 2; i <= n; i++ {
		c := int(s[i-1] - '0')
		l, r = c, l*10+c
		if 1 <= l {
			dp[i] = dp[i-1]
		}
		if r >= 10 && r <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
