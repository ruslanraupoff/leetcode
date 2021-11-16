class Solution:
    def findKthNumber(self, m: int, n: int, k: int) -> int:
        if m > n:
            m, n = n, m

        def count(v, m, n):
            c = 0
            for i in range(1, m+1):
                c += min(v//i, n)
            return c

        l, r = 1, m*n
        while l < r:
            mid = l + (r-l)//2
            c = count(mid, m, n)
            if c >= k:
                r = mid
            else:
                l = mid + 1

        return r
