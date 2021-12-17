class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        m, n = len(matrix), len(matrix[0])
        if m == 0 or n == 0:
            return 0
        

        mxe = 0
        dp = [[0]*(m+1) for _ in range(n+1)]
        
        for i in range(1, m+1):
            for j in range(1, n+1):
                if matrix[i-1][j-1] == '1':
                    dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
                    mxe = max(mxe, dp[i][j])

        return mxe * mxe