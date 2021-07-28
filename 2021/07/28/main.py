class Solution:
    def beautifulArray(self, n: int) -> List[int]:
        r = [1]
        while len(r) < n:
            r = [i * 2 - 1 for i in r] + [i * 2 for i in r]
        return [i for i in r if i <= n]