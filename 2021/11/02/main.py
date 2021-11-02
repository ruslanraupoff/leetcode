class Solution:
    def uniquePathsIII(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dx = [-1, 1, 0, 0]
        dy = [0, 0, -1, 1]
        queue, paths, walked = [], [], 0

        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    walked |= 1 << (i*n+j)
                    queue.append([i, j])
                    paths.append(1<<(i*n+j))
                elif grid[i][j] == 0:
                    walked |= 1 << (i*n+j)
                
        count = 0

        while len(queue) > 0:
            size = len(queue)
            for s in range(size):
                x, y = queue[s][0], queue[s][1]
                path = paths[s]
                for k in range(4):
                    i, j = x+dx[k], y+dy[k]
                    if i < 0 or m <= i or j < 0 or n <= j or path&(1<<(i*n+j)) != 0:
                        continue
                    
                    if grid[i][j] == 0:
                        queue.append([i, j])
                        paths.append(path|(1<<(i*n+j)))
                    elif grid[i][j] == 2:
                        if path == walked:
                            count += 1
            queue = queue[size:]
            paths = paths[size:]

        return count