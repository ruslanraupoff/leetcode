class Solution:
    def grayCode(self, n: int) -> List[int]:
        def gen(r, n, b):
            if n == 0:
                return r
            l = len(r)
            t = [0] * l
            for i in range(l):
                t[l-i-1] = r[i] + b
            r += t
            return gen(r, n-1, b*2)
        return gen([0], n, 1)