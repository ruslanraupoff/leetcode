class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        r, n, m = 0, len(grid), len(grid[0])
        def dfs(i, j):
            if i < 0 or j < 0 or i >= n or j >= m or grid[i][j] == 0:
                return 0
            grid[i][j] = 0
            return 1 + dfs(i-1, j) + dfs(i, j-1) + dfs(i+1, j) + dfs(i, j+1)
        for i in range(n):
            for j in range(m):
                if grid[i][j]:
                    r = max(r, dfs(i, j))
        return r