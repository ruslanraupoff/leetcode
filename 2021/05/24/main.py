class Solution:
    def toLowerCase(self, s: str) -> str:
        r = ''
        for i in range(len(s)):
            c = s[i]
            if 'A' <= c and c <= 'Z':
                c = chr(ord(s[i]) + ord('a') - ord('A'))
            r += c
        return r