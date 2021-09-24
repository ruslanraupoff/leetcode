class Solution:
    def tribonacci(self, n: int) -> int:
        t = [0, 1, 1]
        for i in range(n):
            t.append(t[i] +  t[i+1] +  t[i+2])
        return t[n]