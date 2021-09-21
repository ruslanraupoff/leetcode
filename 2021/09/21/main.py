class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        r, i, cnt = 0, 0, 0
        while i < len(nums):
            while i < len(nums) and nums[i] == 1:
                i += 1
                cnt += 1
            r = max(r, cnt)
            i += 1
            cnt = 0
        return r
