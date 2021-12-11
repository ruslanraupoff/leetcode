class Solution:
    def nthMagicalNumber(self, n: int, a: int, b: int) -> int:
        mod = 10**9 + 7
        if a > b:
            a, b = b, a

        def gcd(a, b):
            if a < b:
                a, b = b, a
            while b != 0:
                a, b = b, a%b
            return a

        def lcm(a, b):
            return a * b // gcd(a, b)

        def magical_of(num, a, b, m):
            return num//a + num//b - num//m

        m = lcm(a, b)
        l, r = a*n // 2, b*n

        while True:
            mid = l + (r - l) // 2
            count = magical_of(mid, a, b, m)
            if count < n:
                l = mid + 1
            elif count > n:
                r = mid - 1
            else:
                res = mid - min(mid % a, mid % b)
                return res % mod

        