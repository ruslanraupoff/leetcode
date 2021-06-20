class Solution:
    def swimInWater(self, grid: List[List[int]]) -> int:
        arr = [(grid[0][0], 0, 0)]
        cc = set()
        res = 0
        heapq.heapify(arr)
        while True:
            n, i, j = heapq.heappop(arr)
            res = max(res, n)
            if i == j == len(grid) - 1:
                return res
            cc.add((i, j))
            for ii, jj in ([i + 1, j], [i - 1, j], [i, j - 1], [i, j + 1]):
                if 0 <= ii < len(grid) and 0 <= jj < len(
                        grid[0]) and (ii, jj) not in cc:
                    heapq.heappush(arr, (grid[ii][jj], ii, jj))
