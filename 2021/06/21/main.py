class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        r = []
        for i in range(numRows):
            l = [1 for _ in range(i+1)]
            for j in range(1, len(l) - 1):
                l[j] = r[i-1][j-1] + r[i-1][j]
            r.append(l)
        return r