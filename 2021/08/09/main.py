class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        r, m = "", 0
        i, j = len(num1) - 1, len(num2) - 1
        while i >= 0 or j >= 0:
            a, b = 0, 0

            if i >= 0:
                a = int(num1[i])
            if j >= 0:
                b = int(num2[j])
            
            s = a + b + m
            m = s // 10
            r = str(s % 10) + r
            i -= 1
            j -= 1
        return '1' + r if m else r 
        