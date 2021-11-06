class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        x = 0
        for n in nums:
            x ^= n

        l, a, b = x & -x, 0, 0

        for n in nums:
            if n & l == 0:
                a ^= n
            else:
                b ^= n
                
        return [a, b]