class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        ps = list(pattern)
        ss = s.split(" ")
        
        if len(ps) != len(ss):
            return False

        def is_match(s1, s2):
            m = {}
            for i in range(len(s1)):
                c = s1[i]
                if c in m:
                    if m[c] != s2[i]:
                        return False
                else:
                    m[c] = s2[i]
            return True

        return is_match(ps, ss) and is_match(ss, ps)
    