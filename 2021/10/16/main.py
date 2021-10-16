class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        l, r = 0, 0
        p, t = 0, 0

        for x in prices:
            l = min(x, l)
            p = max(x - l, p)
            r = min(x - p, r)
            t = max(x - r, t)
        return t