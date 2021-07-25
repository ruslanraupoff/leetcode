class Solution:
    def findIntegers(self, n: int) -> int:
        b = bin(n)[2:][::-1]
        dp = [[1,1] for _ in range(len(b))]
        r = 1 if b[0]=='0' else 2
        for i in range(1, len(b)):
            dp[i][0] = dp[i-1][0] + dp[i-1][1]
            dp[i][1] = dp[i-1][0]
            if b[i-1:i+1] == '01':
                r += dp[i][0]
            elif b[i-1:i+1] == '11':
                r = dp[i][0] + dp[i][1]
        return r