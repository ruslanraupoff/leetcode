class Solution:
    def minFlipsMonoIncr(self, s: str) -> int:
        o = f = 0
        for c in s:
            if c == "0":
                f = min(f + 1, o)
            else:
                o += 1
        return f