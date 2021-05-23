class Solution:
    def shortestSuperstring(self, words: List[str]) -> str:
        n = len(words)
        
        dis = [[0]*n for _ in range(n)]
        for i in range(n):
            for j in range(n):
                a, b = words[i], words[j]
                for k in range(1, len(a)):
                    if b.startswith(a[k:]):
                        dis[i][j] = len(a) - k
                        break
        
        dp = [[''] * n for _ in range(1 << n)]
        for i in range(1 << n):
            for k in range(n):
                if not (i & (1 << k)):
                    continue
                if i == 1 << k:
                    dp[i][k] = words[k]
                    continue
                for j in range(n):
                    if j == k:
                        continue
                    if i & (1 << j):
                        s = dp[i ^ (1 << k)][j]
                        s += words[k][dis[j][k]:]
                        if dp[i][k] == '' or len(s) < len(dp[i][k]):
                            dp[i][k] = s
        r, mn = '', float('inf')
        for i in range(n):
            s = dp[(1 << n) - 1][i]
            if len(s) < mn:
                r, mn = s, len(s)
        return r
