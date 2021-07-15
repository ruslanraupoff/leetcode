class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        m1, m2 = {}, {}
        for i in range(len(s)):
            if m1.get(s[i], '') != m2.get(t[i], ''):
                return False
            m1[s[i]] = i
            m2[t[i]] = i
        return True