class Solution:
    def orderlyQueue(self, s: str, k: int) -> str:
        n = len(s)
        if n == 1:
            return s
        if k > 1:
            return ''.join(sorted(s))
        s2 = 2 * s
        r = "z" * n
        for i in range(n):
            t = s2[i: i + n]
            if t < r:
                r = t
        return r