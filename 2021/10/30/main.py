class Solution:
    def longestDupSubstring(self, s: str) -> str:
        self.res = ""
        def isExist(n):
            seen = set()
            for i in range(len(s)-n+1):
                sub = s[i:i+n]
                if sub in seen:
                    self.res = sub
                    return True
                else:
                    seen.add(sub)
            return False

        l, r = 0, len(s)
        while l < r:
            m = r - (r - l) // 2
            if isExist(m):
                l = m
            else:
                r = m - 1
        return self.res