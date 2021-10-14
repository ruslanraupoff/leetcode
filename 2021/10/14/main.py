class Solution:
    def numSquares(self, n: int) -> int:
        p = [i*i for i in range(1, n+1)]
        dp = [(n + 1) * (n + 1)] * (n+1)
        dp[0] = 0
        for v in p:
            for i in range(v, n+1):
                if dp[i] > dp[i-v]+1:
                    dp[i] = dp[i-v] + 1

        return dp[n]
