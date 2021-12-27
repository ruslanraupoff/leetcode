class Solution:
    def findComplement(self, num: int) -> int:
        r, t = 0, num
        while t > 0:
            t >>= 1
            r <<= 1
            r += 1

        return r ^ num