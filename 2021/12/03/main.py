class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        cur, neg, mx = 1, 1, nums[0]
        for i in range(len(nums)):
            if nums[i] > 0:
                cur, neg = nums[i]*cur, nums[i]*neg
            elif nums[i] < 0:
                cur, neg = nums[i]*neg, nums[i]*cur
            else:
                cur, neg = 0, 1

            mx = max(mx, cur)

            if cur <= 0:
                cur = 1

        return mx