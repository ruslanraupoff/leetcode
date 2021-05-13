class Solution:
    def ambiguousCoordinates(self, s: str):
        def make(s, i, n):
            for j in range(1, n+1):
                left = s[i:i+j]
                right = s[i+j:i+n]
                if (not left.startswith('0') or left == '0') and not right.endswith('0'):
                    yield "".join([left, '.' if right else '', right])

        r = []
        for i in range(1, len(s)-2):
            for x in make(s, 1, i):
                for y in make(s, i+1, len(s)-2-i):
                    r.append("("+x+", "+y+")")
        return r
