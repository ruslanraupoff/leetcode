class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        if sum([len(m) for m in mat]) != (r*c):
            return mat
        d = len(mat[0])
        m = [[0] * c for j in range(r)]
        for i in range(r * c):
            m[i//c][i%c] = mat[i//d][i%d]
        return m