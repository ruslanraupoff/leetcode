class Solution:
    def superpalindromesInRange(self, left: str, right: str) -> int:
        r = 0
        l = [i for i in range(1, 10)]
        def ispalindrome(x):
            return x == x[::-1]
        for i in range(1, 10000):
            l.append(int(str(i) + str(i)[::-1]))
            for j in range(10):
                l.append(int(str(i) + str(j) + str(i)[::-1]))
        for x in l:
            v = x*x
            if int(left) <= v and v <= int(right):
                if ispalindrome(str(v)):
                    r += 1
        return r        
