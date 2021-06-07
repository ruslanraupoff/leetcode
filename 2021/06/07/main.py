class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        a, b = 0, 0
        for i in cost:
            c = i + min(a, b)
            b, a = a, c
        return min(a, b)