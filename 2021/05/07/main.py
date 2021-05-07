class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        r, l1, l2 = 0, len(word1), len(word2)
        dp = [[0 for _ in range(l2 + 1)] for _ in range(l1 + 1)]
        for i in range(1, l1 + 1):
            for j in range(1, l2 + 1):
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
                if word1[i-1] == word2[j-1]:
                    dp[i][j] = dp[i-1][j-1] + 1
        return l1 + l2 - 2 * dp[l1][l2]
