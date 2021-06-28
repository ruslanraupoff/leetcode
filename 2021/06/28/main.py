class Solution:
    def removeDuplicates(self, s: str) -> str:
        l = []
        l.append(s[0])
        for i in range(1, len(s)):
            if len(l) != 0 and l[-1] == s[i]:
                l.pop()
            else:
                l.append(s[i])
        return ''.join(l)