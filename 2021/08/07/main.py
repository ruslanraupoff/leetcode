class Solution:
    def minCut(self, s: str) -> int:
        n = len(s)
        l = [i for i in range(-1, n)]
        dp = [[False]*n for i in range(n)]
        
        for j in range(n):
            for i in range(j, -1, -1):
                if (s[i] == s[j]) and (j-i < 2 or dp[i+1][j-1]):
                    dp[i][j] = True
                    cuts[j+1] = min(l[j+1], l[i]+1)
                    
        return l[-1]