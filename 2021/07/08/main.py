class Solution:
    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        r, n, m = 0, len(nums1), len(nums2)
        dp = [[0] * (m+1) for _ in range(n+1)]
        for i in range(1, n+1):
            for j in range(1, m+1):
                if nums1[i-1] == nums2[j-1]:
                    dp[i][j] = dp[i-1][j-1] + 1
                    r = max(r, dp[i][j])
        return r