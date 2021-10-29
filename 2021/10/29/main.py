class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        dx = [0, 0, -1, 1]
        dy = [-1, 1, 0, 0]
        m, n = len(grid), len(grid[0])

        fresh = 0
        rottens = []
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    fresh += 1
                if grid[i][j] == 2:
                    rottens.append([i, j])
        

        if fresh == 0:
            return 0

        count = -1
        while len(rottens) > 0:
            count += 1
            size = len(rottens)
            for r in range(size):
                x, y = rottens[r][0], rottens[r][1]
                for k in range(4):
                    i, j = x+dx[k], y+dy[k]
                    if i < 0 or m <= i or j < 0 or n <= j or grid[i][j] != 1:
                        continue
                    grid[i][j] = 2
                    fresh -= 1
                    rottens.append([i, j])
            rottens = rottens[size:]

        if fresh > 0:
            return -1
        return count