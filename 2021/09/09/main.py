class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        dp, g, m = [[0]*n for i in range(n)], {tuple(m) for m in mines}, 0
        for i in range(n):
            c = 0
            for j in range(n):
                c = 0 if (i, j) in g else c + 1
                dp[i][j] = c
            c = 0
            for j in range(n-1, -1, -1):
                c = 0 if (i, j) in g else c + 1
                dp[i][j] = min(dp[i][j], c)
        for j in range(n):
            c = 0
            for i in range(n):
                c = 0 if (i, j) in g else c + 1
                dp[i][j] = min(dp[i][j], c)
            c = 0
            for i in range(n-1, -1, -1):
                c = 0 if (i, j) in g else c + 1
                m = max(m, min(dp[i][j], c))
        return m