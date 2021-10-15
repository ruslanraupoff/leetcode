class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        n = len(prices)
        b = [0] * (n+1)
        b[1] = - prices[0]
        s = [0] * (n+1)
        for i in range(2, n+1):
            b[i] = max(b[i-1], s[i-2]-prices[i-1])
            s[i] = max(s[i-1], b[i-1]+prices[i-1])

        return s[n]