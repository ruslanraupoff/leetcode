class Solution:
    def findMin(self, nums: List[int]) -> int:
        n = len(nums)

        if nums[0] < nums[n-1]:
            return nums[0]

        i, j = 0, n-1
        while i < j:
            if nums[i] > nums[i+1]:
                return nums[i+1]

            if nums[j-1] > nums[j]:
                return nums[j]

            i += 1
            j -= 1

        return nums[i]