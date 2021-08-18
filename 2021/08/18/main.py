class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == '0':
            return 0
        n = len(s)
        dp = [0] * (n+1)
        dp[0], dp[1] = 1, 1
        l, r = int(s[0]), 0
        for i in range(2, n+1):
            c = int(s[i-1])
            l, r = c, l*10+c
            if 1 <= l:
                dp[i] = dp[i-1]
            if r >= 10 and r <= 26:
                dp[i] += dp[i-2]

        return dp[n]