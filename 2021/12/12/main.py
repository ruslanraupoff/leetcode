class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        s = sum(nums)

        if s&1 == 1:
            return False

        s = s >> 1
        n = len(nums)

        dp = [[False] * (s+1) for _ in range(n+1)]

        for i in range(n + 1):
            dp[i][0] = True
        for i in range(1, s + 1):
            dp[0][i] = False

        for i in range(1, n + 1):
            for j in range(1, s + 1):
                dp[i][j] = dp[i-1][j]
                if j >= nums[i-1]:
                    dp[i][j] = dp[i][j] or dp[i-1][j-nums[i-1]]

        return dp[n][s]