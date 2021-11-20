class Solution:
    def singleNonDuplicate(self, nums: List[int]) -> int:
        n = len(nums) - 1
        l, r = 0, n-1

        while l < r:
            m = l + (r-l)//2
            if m%2 == 1:
                m -= 1

            if nums[m] != nums[m+1]:
                r = m
            else:
                l = m + 2

        return nums[l]