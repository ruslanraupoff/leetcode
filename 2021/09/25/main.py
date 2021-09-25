class Solution:
    def shortestPath(self, grid: List[List[int]], k: int) -> int:
        n, m = len(grid), len(grid[0])
        s, l, v = 0, [(0, 0, k)], {(0, 0, k)}

        while l:
            ns = []
            for x, y, r in l:
                if x == n - 1 and y == m - 1:
                    return s
                
                for i, j in (x, y+1), (x, y-1), (x+1, y), (x-1, y):
                    if 0 <= i < n and 0 <= j < m:
                        if grid[i][j] == 0:
                            d = (i, j, r)
                            if d not in v:
                                v.add(d)
                                ns.append(d)
                                
                        elif r > 0:
                            d = (i, j, r-1)
                            if d not in v:
                                v.add(d)
                                ns.append(d)
                                
            l = ns
            s += 1
            
        return -1