class Solution:
    def findPaths(self, m: int, n: int, maxMove: int, startRow: int, startColumn: int) -> int:
        dp = [[0] * n for _ in range(m)]
        for k in range(maxMove):
            p = [0] * n
            for i in range(m):
                for j in range(n):
                    r = 0
                    
                    if i == 0:
                        r += 1
                    else:
                        r += p[j]
                    
                    if j == 0:
                        r += 1
                    else:
                        r += p[j-1]
                    
                    if i + 1 == m:
                        r += 1
                    else:
                        r += dp[i+1][j]
                        
                    if j + 1 == n:
                        r += 1
                    else:
                        r += dp[i][j+1]
                        
                    r %= 10 ** 9 + 7
                    p[j] = dp[i][j]
                    dp[i][j] = r
        return dp[startRow][startColumn]