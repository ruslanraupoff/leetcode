class Solution:
    def findMin(self, nums: List[int]) -> int:
        i, n = 1, len(nums)
        while i < n and nums[i-1] < nums[i]:
            i += 1
        return nums[i%n]