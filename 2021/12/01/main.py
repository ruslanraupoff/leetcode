class Solution:
    def rob(self, nums: List[int]) -> int:
        a, b = 0, 0
        for i, v in enumerate(nums):
            if i%2 == 0:
                a = max(a+v, b)
            else:
                b = max(a, b+v)

        return max(a, b)