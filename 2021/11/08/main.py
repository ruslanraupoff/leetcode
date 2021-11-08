class Solution:
    def numTrees(self, n: int) -> int:
        if n < 2:
            return 1
        if n == 2:
            return 2
        if n == 3:
            return 5
        r, d = 0, n // 2 + 1
        for i in range(1, d):
            r += self.numTrees(i-1) * self.numTrees(n-i)
        r *= 2
        if n%2 == 1:
            t = self.numTrees(n // 2)
            r += t * t
        return r