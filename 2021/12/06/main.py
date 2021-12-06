class Solution:
    def minCostToMoveChips(self, position: List[int]) -> int:
        o, e = 0, 0
        for p in position:
            if p % 2:
                e += 1
            else:
                o += 1
        return min(o, e)