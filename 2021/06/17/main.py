class Solution:
    def numSubarrayBoundedMax(self, nums: List[int], left: int, right: int) -> int:
        p, r, n = -1, 0, 0
        for i, num in enumerate(nums):
            if left <= num <= right:
                n = i - p
            if num > right:
                n, p = 0, i
            r += n
        return r