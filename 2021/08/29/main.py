class Solution:
    def minPatches(self, nums: List[int], n: int) -> int:
        i, mx, r = 0, 1, 0
        while mx <= n:
            if i < len(nums) and nums[i] <= mx:
                mx += nums[i]
                i += 1    
            else:
                mx <<= 1
                r += 1
        return r