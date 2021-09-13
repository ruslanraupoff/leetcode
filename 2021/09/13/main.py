class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        p = {c: 0 for c in "balloon"}
        for t in text:
            if t in "balloon":
                p[t] += 1
        mn = min(p['b'], min(p['a'], p['n']))
        mn = min(mn, min(p['l']//2, p['o']//2))
        return mn