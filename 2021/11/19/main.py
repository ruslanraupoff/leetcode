class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        d = 0
        v = x ^ y
        while v > 0:
            v &= v-1
            d += 1
        return d
