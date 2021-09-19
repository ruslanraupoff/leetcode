class Solution:
    def numDistinct(self, s: str, t: str) -> int:
        m, n = len(s), len(t)

        p, dp = 0, [1] * (m + 1)

        for j in range(n):
            dp[j], p = 0, dp[j]
            for i in range(j+1, m + 1):
                if t[j] == s[i-1]:
                    dp[i], p = dp[i-1]+p, dp[i]
                else:
                    dp[i], p = dp[i-1], dp[i]

        return dp[m]