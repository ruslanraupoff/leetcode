class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return len(nums)
        nums.sort()
        mx, r = 0, 1
        for i in range(1, len(nums)):
            if nums[i-1] == nums[i] - 1:
                r += 1
            elif nums[i] != nums[i-1]:
                r = 1
            mx = max(mx, r)
        return mx


