class Solution:
    def candy(self, ratings: List[int]) -> int:
        a, n = 0, len(ratings)
        l, r = [1] * n, [1] * n
        for i in range(1, n):
            if ratings[i] > ratings[i-1]:
                l[i] = l[i-1] + 1
            if ratings[n-i-1] > ratings[n-i]:
                r[n-i-1] = r[n-i] + 1
        for i in range(n):
            a += max(l[i], r[i])
        return a