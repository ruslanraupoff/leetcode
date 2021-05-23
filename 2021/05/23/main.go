package main

import (
	"strings"
)

func shortestSuperstring(words []string) string {
	n, r := len(words), ""
	dis := make([][]int, n)
	for i := 0; i < n; i++ {
		dis[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dis[i][j] = 0
			a, b := words[i], words[j]
			for k := 1; k < len(a); k++ {
				if strings.HasPrefix(b, a[k:]) {
					dis[i][j] = len(a) - k
					break
				}
			}
		}
	}

	dp := make([][]string, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]string, n)
		for k := 0; k < n; k++ {
			if i&(1<<k) == 0 {
				continue
			}
			if i == (1 << k) {
				dp[i][k] = words[k]
				continue
			}
			for j := 0; j < n; j++ {
				if j == k {
					continue
				}
				if i&(1<<j) != 0 {
					s := dp[i^(1<<k)][j]
					s += words[k][dis[j][k]:]
					if dp[i][k] == "" || len(s) < len(dp[i][k]) {
						dp[i][k] = s
					}
				}
			}
		}
	}
	mn := int(^uint(0) >> 1)
	for i := 0; i < n; i++ {
		s := dp[(1<<n)-1][i]
		if len(s) < mn {
			r, mn = s, len(s)
		}
	}
	return r
}
