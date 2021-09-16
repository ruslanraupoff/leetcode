class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        d = [[1, 0], [0, -1], [-1, 0], [0, 1]]
        r, m, n = matrix[0], len(matrix), len(matrix[0])
        i, j, p = (m - 1) * n, 0, [0, n -1]
        while i > 0:
            for _ in range(1, m):
                i -= 1
                p[0] += d[j][0]
                p[1] += d[j][1]
                r.append(matrix[p[0]][p[1]])
            m -= 1
            m, n = n, m
            j = (j + 1) % 4
        return r
