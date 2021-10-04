class Solution:
    def canJump(self, nums: List[int]) -> bool:
        r = len(nums) - 1
        
        for i in range(len(nums)-1, -1, -1):
            if i + nums[i] >= r:
                r = i
        
        return r == 0