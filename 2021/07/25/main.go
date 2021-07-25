package main

func findIntegers(n int) int {
	b := make([]byte, 0, 32)
	for n > 0 {
		b = append(b, byte(n%2)+'0')
		n /= 2
	}

	var dp [][]int
	for i := 0; i < len(b); i++ {
		dp = append(dp, []int{1, 1})
	}
	r := 2
	if b[0] == '0' {
		r = 1
	}
	for i := 1; i < len(b); i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
		bin := b[i-1 : i+1]
		if bin[0] == '0' && bin[1] == '1' {
			r += dp[i][0]
		} else if bin[0] == '1' && bin[1] == '1' {
			r = dp[i][0] + dp[i][1]
		}
	}

	return r
}
