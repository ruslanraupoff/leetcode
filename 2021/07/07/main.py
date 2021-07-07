class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        l = [-10**9-1]
        for i in matrix:
            l += i
        l.sort()
        return l[k]
