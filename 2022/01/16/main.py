class Solution:
    def maxDistToClosest(self, seats: List[int]) -> int:
        d, e = 0, 0
        for i in range(len(seats)):
            if e == i:
                d = e
            else:
                d = max(d, (e+e%2)/2)
            if seats[i] == 1:
                e = 0
            else:
                e += 1

        return max(d, e)