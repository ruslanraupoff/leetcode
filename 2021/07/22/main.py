class Solution:
    def partitionDisjoint(self, nums: List[int]) -> int:
        p, n, mx = 0, nums[0], nums[0]
        for i in range(len(nums)):
            mx = max(mx, nums[i])
            if nums[i] < n:
                p = i
                n = mx
        return p + 1

