class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        l = []
        for t in tokens:
            if t in "+-*/":
                a, b = l[-2],  l[-1]
                l = l[:-2]
                if t == '+':
                    l.append(a+b)
                if t == '-':
                    l.append(a-b)
                if t == '*':
                    l.append(a*b)
                if t == '/':
                    l.append(int(a/b))
            else:
                l.append(int(t))
        return l[0]
