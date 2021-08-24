class Solution:
    def complexNumberMultiply(self, num1: str, num2: str) -> str:
        def parse(y):
            z = y.split('+')
            return int(z[0]), int(z[1][:-1])
        a, b = parse(num1)
        c, d = parse(num2)
        return str(a*c-b*d) + '+' + str(a*d+b*c) + 'i'
        