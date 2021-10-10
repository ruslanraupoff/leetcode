class Solution:
    def rangeBitwiseAnd(self, left: int, right: int) -> int:
        r = 0
        while left >= 1 and right >=1:
            if left == right:
                return left << r
            left >>= 1
            right >>= 1
            r += 1

        return 0