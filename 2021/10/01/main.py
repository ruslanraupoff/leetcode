class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        m, n = len(text1), len(text2)
        cur = [0]*(m+1)
        prev = cur
        for j in range(n):
            for i in range(m):
                rec = max(prev[i+1], cur[i])
                if text1[i] == text2[j]:
                    cur[i+1] = max(rec, prev[i]+1)
                else:
                    cur[i+1] = rec
            cur, prev = prev, cur
        return prev[m]