class Solution:
    def countPrimes(self, n: int) -> int:
        primes = [True for _ in range(n+1)]
        i = 2
        for i in range(2, n+1):
            if primes[i] == False:
                continue
            for j in range(i*i, n, i):
                primes[j] = False
        r = 0
        for i in range(2, n):
            if primes[i]:
                r += 1
        return r
