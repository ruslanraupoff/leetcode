class Solution:
    def stoneGameVII(self, stones: List[int]) -> int:
        d = len(stones)
        dp = [[0] * d for _ in range(d)]
        for i in range(d-1, -1, -1):
            sm = stones[i]
            for j in range(i+1, d):
                sm += stones[j]
                dp[i][j] = max(sm-stones[i]-dp[i+1][j], sm-stones[j]-dp[i][j-1])
        return dp[0][-1]
