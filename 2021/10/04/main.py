class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dx = [-1, 1, 0, 0]
        dy = [0, 0, -1, 1]

        r = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 0:
                    continue
                r += 4
                for k in range(4):
                    x = i + dx[k]
                    y = j + dy[k]
                    if 0 <= x and x < m and 0 <= y and y < n and grid[x][y] == 1:
                        r -= 1
        return r