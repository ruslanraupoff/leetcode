class Solution:
    def carPooling(self, trips: List[List[int]], capacity: int) -> bool:
        l = {}
        for t in trips:
            num, start, end = t[0], t[1], t[2]
            if start in l:
                l[start] += num
            else:
                l[start] = num
            if end in l:
                l[end] -= num
            else:
                l[end] = num

        p = 0
        for c in l:
            p += c
            if p > capacity:
                return False
        return True