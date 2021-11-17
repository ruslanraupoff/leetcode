class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [1] * (n+1)

        for i in range(2, m+1):
            for j in range(2, n+1):
                dp[j] += dp[j-1]

        return dp[n]