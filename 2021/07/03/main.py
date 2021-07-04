class Solution:
    def maxSumSubmatrix(self, matrix: List[List[int]], k: int) -> int:
        m, n = len(matrix), len(matrix[0])
        for x in range(m):
            for y in range(n - 1):
                matrix[x][y+1] += matrix[x][y]
        res, maxSum = 0, -float("Inf")
        for y1 in range(n):
            for y2 in range(y1, n):
                bilist = [0]
                s = 0
                for x in range(m):
                    s += matrix[x][y2] - (matrix[x][y1-1] if y1 > 0 else 0)
                    i = bisect.bisect_left(bilist, s - k)
                    if i < len(bilist):
                        maxSum = max(maxSum, s - bilist[i])
                    bisect.insort(bilist, s)
        return maxSum
