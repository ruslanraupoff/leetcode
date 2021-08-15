class Solution:
    def minWindow(self, s: str, t: str) -> str:
        a = [0 for _ in range(58)]
        for c in t:
            k = ord(c) - 65
            a[k] = a[k] + 1
        l, r, mx = 0, 0, 10**6
        i, j, n = 0, -1, len(t)
        while r < len(s):
            k = ord(s[r]) - 65
            a[k] -= 1
            if a[k] >= 0: n -= 1
            while n == 0:
                m = r - l + 1
                if m < mx:
                    i, j, mx = l, r, m
                k = ord(s[l]) - 65
                a[k] += 1
                if a[k] > 0: n += 1
                l += 1
            r += 1
        return s[i: j + 1]