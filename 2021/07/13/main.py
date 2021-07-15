class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        l, r = -1, len(nums)
        m = l + (r - l) // 2
        while l + 1 <= m and m < r - 1:
            if nums[m] < nums[m+1]:
                l = m
            else:
                r = m + 1
            m = l + (r - l) // 2
        return l + 1