class Solution:
    def numDecodings(self, s: str) -> int:
        m = 10**9 + 7
        l, r = 1, 1
        if s[0] == '*':
            r = 9
        elif s[0]=='0':
            r = 0
        for i in range(1, len(s)):
            t = r
            if s[i] == '*':
                r = 9*r % m
                if s[i-1]=='1':
                    r = (r + 9*l) % m
                elif s[i-1] == '2':
                    r = (r + 6*l) % m
                elif s[i-1] == '*':
                    r = (r + 15*l) % m
            else:
                if s[i] == '0':
                    r = 0
                if s[i-1]=='1':
                    r = (r + l) % m
                elif s[i-1] == '2' and s[i] <= '6':
                    r = (r + l) % m
                elif s[i-1] == '*':
                    r = (r + (2 if s[i] <= '6' else 1) * l) % m
            l = t
        return r