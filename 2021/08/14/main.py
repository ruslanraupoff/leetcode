class Solution:
    def removeBoxes(self, boxes: List[int]) -> int:
        @lru_cache(None)
        def dfs(l: int, r: int, k: int) -> int:
            if l > r:
                return 0
            while r > l and boxes[r] == boxes[r-1]:
                r -= 1
                k += 1
            res = dfs(l, r-1, 0) + (k+1)**2
            for i in range(l, r):
                if boxes[i] == boxes[r]:
                    res = max(res, dfs(l, i, k+1) + dfs(i+1, r-1, 0))
            return res

        return dfs(0, len(boxes) - 1, 0)
