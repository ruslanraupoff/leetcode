class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        i, j, k = 0, 0, len(nums)-1
        while j <= k:
            n = nums[j]
            if n == 0:
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j += 1
            elif n == 1:
                j += 1
            elif n == 2:
                nums[j], nums[k] = nums[k], nums[j]
                k -= 1
        return nums