class Solution:
    def matrixRankTransform(self, matrix: List[List[int]]) -> List[List[int]]:
        n, m = len(matrix), len(matrix[0])
        r = [0] * (m + n)

        d = defaultdict(list)
        for i in range(n):
            for j in range(m):
                d[matrix[i][j]].append([i, j])

        def find(i):
            if p[i] != i:
                p[i] = find(p[i])
            return p[i]
        
        for val in sorted(d):
            p = list(range(m + n))
            t = r[:]
            for i, j in d[val]:
                i, j = find(i), find(j + n)
                p[i] = j
                t[j] = max(t[i], t[j])
            
            for i, j in d[val]:
                r[i] = r[j + n] = matrix[i][j] = t[find(i)] + 1
        return matrix