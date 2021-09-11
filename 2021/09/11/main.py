class Solution:
    def calculate(self, s: str) -> int:
        num, res, sign, stack = 0, 0, 1, []

        for i in range(len(s)):
            c = s[i]
            if c in '0123456789':
                num = 0
                while i < len(s) and c >= '0' and c <= '9':
                    num = 10 * num + ord(c) - ord('0')
                    i += 1
                res += sign * num
                i -= 1
            if c == '+':
                sign = 1
            if c == '-':
                sign = -1
            if c == '(':
                stack.append(res, sign)
                res = 0
                sign = 1
            if c == ')':
                last = stack.pop()
                sign = last[1]
                res = sign * res + last[0]
                sign = 1

        return res