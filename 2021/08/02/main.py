class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hs = {}
        for i in range(len(nums)):
            n = nums[i]
            if n in hs:
                return [hs[n], i]
            hs[target - n] = i
        
