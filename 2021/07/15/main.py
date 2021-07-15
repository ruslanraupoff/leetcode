class Solution:
    def triangleNumber(self, nums: List[int]) -> int:
        c, n =  0, len(nums)
        nums.sort()
        for i in range(n-1, 0, -1):
            l, r = 0, i-1
            while l < r:
                if nums[l] + nums[r] > nums[i]:
                    c += r - l
                    r -= 1
                else:
                    l += 1
        return c