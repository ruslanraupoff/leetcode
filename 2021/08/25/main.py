class Solution:
    def judgeSquareSum(self, c: int) -> bool:
        i = 2
        while i * i <= c:
            cnt = 0
            if c % i == 0:
                while c % i == 0:
                    cnt += 1
                    c /= i
                if i % 4 == 3 and cnt % 2 != 0:
                    return False
            i += 1
        return c % 4 != 3
