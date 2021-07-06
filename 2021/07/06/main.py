class Solution:
    def minSetSize(self, arr: List[int]) -> int:
        c, s, m = 0, {}, len(arr) / 2
        for a in arr:
            if a in s:
                s[a] += 1
            else:
                s[a] = 1
        s = sorted(s.items(), key=lambda x: -x[1])
        for i, p in enumerate(s):
            c += p[1]
            if c >= m:
                return i + 1