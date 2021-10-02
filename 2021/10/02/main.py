class Solution:
    def calculateMinimumHP(self, dungeon: List[List[int]]) -> int:
        n, m = len(dungeon[0]), len(dungeon)
        if m == 0 or n == 0:
            return 1
        dp = [[1<<63 - 1] * (n+1) for _ in range(n+1)]
        dp[m][n-1] = 1
        for i in range(m-1, -1, -1):
            for j in range(n-1, -1, -1):
                mn = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
                dp[i][j] = max(mn, 1)

        return dp[0][0]