class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        s, mn = 0, 0
        for n in nums:
            s += n
            mn = min(s, mn)
        return 1 - mn