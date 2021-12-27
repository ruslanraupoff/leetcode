class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        def dist(x):
            return x[0] * x[0] + x[1] * x[1]

        return heapq.nsmallest(k, points, key=lambda x: dist(x))