class Solution:
    def calculate(self, s: str) -> int:
        def operate(a, b, opt):
            if opt == '+':
                return a + b
            if opt == '-':
                return a - b
            if opt == '*':
                return a * b
            return a / b
        x, y, z = 0, 0, 0
        c = '+'

        self.idx = 0
        def nextN():
            n = 0
            while self.idx < len(s) and s[self.idx] == ' ':
                self.idx += 1
            
            while self.idx < len(s) and '0' <= s[self.idx] and s[self.idx] <= '9':
                n = 10 * n + ord(s[self.idx]) - ord('0')
                self.idx += 1
            return n

        def nextOpt():
            opt = '+'
            while self.idx < len(s) and s[self.idx] == ' ':
                self.idx =+ 1
            if self.idx < len(s):
                opt = s[self.idx]
                self.idx += 1
            return opt

        y = nextN()
        while self.idx < len(s):
            d = nextOpt()
            z = nextN()

            if d == '*' or d == '/':
                y = operate(y, z, d)
            else:
                x = operate(x, y, c)
                c = d
                y = z
        
        return operate(x, y, c)