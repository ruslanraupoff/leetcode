class Solution:
    def arrangeCoins(self, n: int) -> int:
        return math.floor((-1.0+math.sqrt(1+8*n)) * 0.5)