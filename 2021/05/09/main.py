class Solution:
    def isPossible(self, target: List[int]) -> bool:
        s, l = sum(target), len(target)
        target.sort()
        if l == 1:
            return target[0] == 1
        while True:
            a = target[l-1]
            s -= a
            if a == 1 or s == 1:
                return True
            if a < s or a % s == 0:
                return False
            a %= s
            s += a
            target[l-1] = a
            target.sort()
