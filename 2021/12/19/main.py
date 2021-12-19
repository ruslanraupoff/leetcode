class Solution:
    def decodeString(self, s: str) -> str:
        def dfs(i):
            c = r = ''
            while i < len(s):
                if s[i] == '[':
                    t, i = dfs(i+1)
                    r += int(c) * t
                    c = ''
                elif s[i] == ']':
                    return r, i
                elif '0' <= s[i] <= '9':
                    c += s[i]
                else:
                    r += s[i]
                i += 1
            return r
        return dfs(0)
