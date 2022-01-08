class Solution:
    def cherryPickup(self, grid: List[List[int]]) -> int:
        dp = [[[0]*71 for i in range(71)] for i in range(71)]
        m, n = len(grid), len(grid[0])
        for i in range(m-1, -1, -1):
            for j in range(n):
                for k in range(n):
                    r = grid[i][j]
                    if j != k:
                        r += grid[i][k]
                    if i != m-1:
                        mx = 0
                        for x in range(j-1, j+2):
                            for y in range(k-1, k+2):
                                if x >= 0 and x < n and y >= 0 and y < n:
                                    mx = max(mx, dp[i+1][x][y])
                        r += mx

                    dp[i][j][k] = r
        
        return dp[0][0][n-1]

