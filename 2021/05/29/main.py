class Solution:
    def totalNQueens(self, n: int) -> int:
        self.s = 0
        def dfs(r, c, p, a):
            if r == n:
                self.s += 1
                return
            bits = ~(c|p|a) & ((1<<n) - 1)
            while bits:
                pos = bits & (-bits)
                bits &= (bits - 1)
                dfs(r+1, c|pos, (p|pos)<<1, (a|pos)>>1)
        dfs(0, 0, 0, 0)
        return self.s