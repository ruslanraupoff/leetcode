class Solution:
    def numberOfArithmeticSlices(self, nums: List[int]) -> int:
        n = len(nums)
        dp = [collections.defaultdict(int) for i in range(n)]

        res = 0
        for i in range(1, n):
            for j in range(i):
                d = nums[i] - nums[j]
                dp[i][d] += dp[j][d] + 1
                if dp[j][d]:
                    res += dp[j][d]
        return res